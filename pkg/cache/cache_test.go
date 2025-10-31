package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

// 主键缓存sql
type PrimaryKeySqlCache struct {
	DBClient *SqlClient
}

func (p *PrimaryKeySqlCache) CacheGet(rid string, tableName string, dbData interface{}, cacheCol string, keys ...interface{}) (interface{}, bool, error) {
	session := p.DBClient.ReadEngine().Table(tableName).Where("rid = ?", rid)
	// 如果没有特定的列, 则缓存整个表
	if cacheCol != "" {
		session.Cols(cacheCol)
	}
	find, err := session.Get(dbData)
	return dbData, find, err
}

func (p *PrimaryKeySqlCache) CacheUpdate(rid string, tableName string, dbData interface{}, cacheCol string, keys ...interface{}) error {
	session := p.DBClient.Table(tableName).Where("rid = ?", rid)
	if cacheCol != "" {
		session.Cols(cacheCol)
	}
	_, err := session.Update(dbData)
	return err

}

func (p *PrimaryKeySqlCache) CacheInsert(rid string, tableName string, dbData interface{}) error {
	aff, err := p.DBClient.Table(tableName).Insert(dbData)
	if aff != 1 {
		return fmt.Errorf("rid:%v cache:%v insert aff%v-1", rid, tableName, aff)
	}
	return err
}



type SqlClient struct {
	*xorm.EngineGroup
}

func (d *SqlClient) ReadEngine() *xorm.Engine {
	// 如果没有从库就用主库读
	readEngines := d.Slaves()
	if len(readEngines) == 0 {
		return d.Engine
	}
	// 随机选取一个从库
	return readEngines[rand.Intn(len(readEngines))]
}

func parseParams2Dsn(user, passwd, addr, db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, passwd, addr, db)
}

func newMysqlClent(mysqlConfig DBMSConfig) (*SqlClient, error) {
	databaseDsns := make([]string, 0, len(mysqlConfig.Slaves)+1)
	masterDsn := parseParams2Dsn(mysqlConfig.Master.User, mysqlConfig.Master.Passwd, mysqlConfig.Master.Addr, mysqlConfig.Master.DB)
	databaseDsns = append(databaseDsns, masterDsn)
	for _, slave := range mysqlConfig.Slaves {
		slavesDsn := parseParams2Dsn(slave.User, slave.Passwd, slave.Addr, slave.DB)
		databaseDsns = append(databaseDsns, slavesDsn)
	}
	engineGroup, err := xorm.NewEngineGroup("mysql", databaseDsns)
	if err != nil {
		return nil, err
	}
	engineGroup.SetMaxIdleConns(int(mysqlConfig.MaxIdleConn))
	engineGroup.SetMaxOpenConns(int(mysqlConfig.MaxOpenConn))
	engineGroup.AddHook(&MysqlHook{})
	return &SqlClient{engineGroup}, nil
}

var dbMutex sync.Mutex
var dbClientMap = make(map[string]*SqlClient)

type DBConfig struct {
	Addr   string `yaml:"addr"`
	DB     string `yaml:"db"`     // 数据库
	User   string `yaml:"user"`   // 数据库用户
	Passwd string `yaml:"passwd"` // 数据库密码
}

type DBMSConfig struct {
	Tag         string      `yaml:"tag"`    // 用来分库
	Master      DBConfig    `yaml:"master"` // 主库
	Slaves      []*DBConfig `yaml:"slaves"` // 从库
	MaxOpenConn int         `yaml:"max_open_conn"`
	MaxIdleConn int         `yaml:"max_idle_conn"`
}

func getOperationType(sql string) string {
	if len(sql) < 6 {
		return "OTHER"
	}
	return strings.ToUpper(sql[:6])
}

// mysql的读写钩子
type MysqlHook struct {
}

func (h *MysqlHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return c.Ctx, nil
}

func (h *MysqlHook) AfterProcess(c *contexts.ContextHook) error {
	// 获取操作类型 (SELECT, INSERT, UPDATE, DELETE)
	opType := getOperationType(c.SQL)
	fmt.Printf("[xorm]:%v use time:%vms - sql: %v args[%v]\n", opType, c.ExecuteTime.Milliseconds(), c.SQL, c.Args)
	return nil
}

func NewSqlClent(mysqlConfigs []DBMSConfig, tag string) (*SqlClient, error) {
	// 初始化数据库xorm
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if dbClientMap[tag] != nil {
		return dbClientMap[tag], nil
	}

	for _, mysqlConfig := range mysqlConfigs {
		if mysqlConfig.Tag == tag {
			client, err := newMysqlClent(mysqlConfig)
			if err != nil {
				return nil, err
			}
			dbClientMap[tag] = client
			return client, nil
		}
	}
	strJson, _ := json.Marshal(mysqlConfigs)
	errStr := fmt.Sprintf("mysql tag:%v url[%v] db New Engine is not found", tag, string(strJson))
	return nil, fmt.Errorf(errStr)
}

func TestIntCache(t *testing.T) {
	// 下面是示例, 实际代码会不同
	hanlder := NewHandleCache(CacheConfig{
		TableName:   (&dbmodel.RoleT{}).SubName(),
		ColName:     dbmodel.TRole.Base,
		Conn:        redisClient.Client,
		CacheLock:   false,
		Stats:       xcache.NewCacheStat(time.Second*2, func(hit, miss uint64) { redisClient.CacheHitRate(dbmodel.TRole.Base, hit, miss) }),
		Encrypt:     &xcache.JsonType{Val: &db_extra.RoleItemT{}},
		TableStruct: dbmodel.RoleT{},
		Manger:      &xclient.PrimaryKeySqlCache{DBClient: db},
	})
	hanlder.Get()
	hanlder.Update()
	hanlder.Insert()
}
