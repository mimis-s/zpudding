package cache

import (
	"reflect"

	"github.com/go-redis/redis/v8"
)

/*
并发问题:
1: 在load数据还没有读取到redis中的时候, 这个时候外部数据更新了, 要update
1解决: 所有操作都应该给这条redis记录加一个乐观锁, update的时候发现这条记录已经被load锁住了, 就会等待load结束之后再操作
*/
type UpdateFuncHandle func(rid string, data interface{}, keys ...interface{}) error

type GetFuncHanel func(rid string, keys ...interface{}) (interface{}, bool, error)

type CacheInfo interface {
	// CacheBackSave(rid string, keys ...interface{}) error                                                                 // 将缓存数据读入数据库, 删除缓存
	Get(rid string, keys ...interface{}) (interface{}, bool, error)                                                      // 获取cache数据
	Update(rid string, updateData func(data interface{}) (interface{}, error), keys ...interface{}) (interface{}, error) // 更新cache数据
	Insert(rid string, data interface{}, keys ...interface{}) error                                                      // 插入cache数据
}

type CacheConfig struct {
	TableName   string // key的前缀
	ColName     string // 缓存数据的名字
	Conn        *redis.Client
	CacheLock   bool        // 操作缓存的时候加分布式锁
	Stats       *CacheStat  // 缓存命中率统计
	Encrypt     Serialize   // 编解码方式(如果客户端不传则为string)
	Lifetime    int         // 缓存存活时间(s)
	TableStruct interface{} // 要缓存的表
	Manger      CacheManger
}

func NewHandleCache(cacheConfig CacheConfig) CacheInfo {
	if cacheConfig.Manger == nil {
		return nil
	}
	if cacheConfig.Lifetime <= 0 {
		// 默认存活60s
		cacheConfig.Lifetime = 60
	}
	return &HandleCacheInfo{
		tableName: cacheConfig.TableName,
		colName:   cacheConfig.ColName,
		tableType: reflect.TypeOf(cacheConfig.TableStruct),
		manger:    cacheConfig.Manger,
		cache: &cacheBase{
			Conn:      cacheConfig.Conn,
			CacheLock: cacheConfig.CacheLock,
			Stats:     cacheConfig.Stats,
			Encrypt:   cacheConfig.Encrypt,
			Lifetime:  cacheConfig.Lifetime,
		},
	}
}
