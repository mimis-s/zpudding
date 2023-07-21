package dialect

import (
	"fmt"
	"os"
	"testing"

	"xorm.io/xorm"
	"xorm.io/xorm/dialects"
	"xorm.io/xorm/log"
)

func TestMysql(t *testing.T) {
	orm := &xorm.Engine{}
	dataSourceName := "dev" + ":" + "dev123" + "@tcp(" + "192.168.1.22" + ")/" + "idle-likun" + "?charset=utf8"
	// dialects.RegisterDriver("mysql", &mysqlDriver{})
	// 修复了sync时，结构体包含json的字段时，sync到库里的字段是text类型
	dialects.RegisterDialect("mysql", func() dialects.Dialect { return &Mysql{} })

	orm, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("new xorm engine[%v] error:%v\n", dataSourceName, err)
		os.Exit(1)
	}
	orm.SetMaxOpenConns(64)
	orm.SetMaxIdleConns(64)
	err = orm.Ping()
	if err != nil {
		fmt.Printf("ping database[%v] error:%v\n", dataSourceName, err)
		os.Exit(1)
	}

	orm.ShowSQL(true)
	orm.SetLogLevel(log.LOG_WARNING)

	// err = orm.Sync2()
	// if err != nil {
	// 	panic(err)
	// }
}
