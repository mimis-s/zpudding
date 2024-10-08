package cache

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

var (
	prefixCache = "cache.data.int."
)

type HandleCacheInfo struct {
	ServiceName string // 服务名字
	CacheName   string // 缓存数据的名字
	Conn        *redis.Client
	Engine      *xorm.Engine
	Stats       *CacheStat       // 缓存命中率统计
	Encrypt     Serialize        // 编解码方式(如果客户端不传则为string)
	GetFunc     GetFuncHanel     // db的读写函数
	UpdateFunc  UpdateFuncHandle // db的读写函数

	/*
		在logout的时候, 应该把所有存储在redis的数据都更新到mysql里面, 但是同一条数据, 可能因为查询id不一样, 导致相同的数据被存入不同的redis key中
		这样的数据即使更新到数据库里, 它的一致性也无法保证了, 这样的话, 我们在定义key的时候就不能简单的用查询的key拼接使用了
		比如角色信息, 我可以通过roleid查询, 也可以通过playername查询, 但是这两个key查询出来的东西都是一样的
		解决方案: 在get函数里面返回当前查询出来字段的唯一key(主键/mysql唯一索引, 后面还可以在update发挥用处), 用来作为redis的key, 同时记录这个key, 等到logout的时候, 外部不需要传入key, 内部将所有查询出来的key回写sql
	*/
}

func judgeKeyType(key interface{}) (string, error) {
	switch key.(type) {
	case int:
		return strconv.Itoa(key.(int)), nil
	case int32:
		return strconv.Itoa(int(key.(int32))), nil
	case int64:
		return strconv.FormatInt(key.(int64), 10), nil
	case string:
		return key.(string), nil
	default:
		return "", fmt.Errorf("key[%v] type[%v] is not found", key, reflect.TypeOf(key))
	}
}

// func (s *HandleCacheInfo) LoadToCache(id int64, datas interface{}, keys ...interface{}) (interface{}, error) {
// 	strID := strconv.FormatInt(id, 10)
// 	for _, key := range keys {
// 		strKey, err := judgeKeyType(key)
// 		if err != nil {
// 			return nil, err
// 		}
// 		strID += "." + strKey
// 	}

// 	prefixCacheForID := prefixCache + s.ServiceName + "." + s.CacheName

// 	if s.GetFunc == nil || s.UpdateFunc == nil {
// 		return nil, fmt.Errorf("try get id(%v) data cache[%v] db func is nil", strID, prefixCacheForID)
// 	}

// 	reData, err := LockGetOrInsertCache(s, prefixCacheForID, strID, func() (interface{}, error) {
// 		var dataCache interface{}
// 		var err error
// 		if datas != "" {
// 			dataCache = datas
// 		} else {
// 			dataCache, err = s.GetFunc(id, keys)
// 			if err != nil {
// 				return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
// 			}
// 		}
// 		return dataCache, nil
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("try set id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
// 	}

// 	return reData, nil
// }

func (s *HandleCacheInfo) LogoutCache(id int64, keys ...interface{}) error {
	strID := strconv.FormatInt(id, 10)
	for _, key := range keys {
		strKey, err := judgeKeyType(key)
		if err != nil {
			return err
		}
		strID += "." + strKey
	}

	prefixCacheForID := prefixCache + s.ServiceName + "." + s.CacheName

	if s.GetFunc == nil || s.UpdateFunc == nil {
		return fmt.Errorf("try get id(%v) data cache[%v] db func is nil", strID, prefixCacheForID)
	}

	LockDelCache(s, prefixCacheForID, strID, func(data interface{}) error {
		err := s.UpdateFunc(id, keys, data)
		if err != nil {
			return fmt.Errorf("try get and update id(%v) data[%v] cache[%v] is err:%v", strID, data, prefixCacheForID, err)
		}
		return nil
	})

	return nil
}

func (s *HandleCacheInfo) Get(id int64, keys ...interface{}) (interface{}, bool, error) {
	strID := strconv.FormatInt(id, 10)
	for _, key := range keys {
		strKey, err := judgeKeyType(key)
		if err != nil {
			return nil, false, err
		}
		strID += "." + strKey
	}

	prefixCacheForID := prefixCache + s.ServiceName + "." + s.CacheName

	if s.GetFunc == nil || s.UpdateFunc == nil {
		return nil, false, fmt.Errorf("try get id(%v) data cache[%v] db func is nil", strID, prefixCacheForID)
	}

	data, bFind, err := LockGetCache(s, prefixCacheForID, strID, func() (interface{}, bool, error) {

		dataCache, find, err := s.GetFunc(id, keys)
		if err != nil {
			return "", false, fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
		}
		return dataCache, find, nil
	})
	if err != nil {
		return "", bFind, fmt.Errorf("try get and insert id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
	}

	return data, bFind, nil
}

// 获取数据, 如果数据不存在, 则插入之后获取
func (s *HandleCacheInfo) GetOrInsert(id int64, dbInsertFunc func() (interface{}, error), keys ...interface{}) (interface{}, error) {
	strID := strconv.FormatInt(id, 10)
	for _, key := range keys {
		strKey, err := judgeKeyType(key)
		if err != nil {
			return nil, err
		}
		strID += "." + strKey
	}

	prefixCacheForID := prefixCache + s.ServiceName + "." + s.CacheName

	if s.GetFunc == nil || s.UpdateFunc == nil {
		return nil, fmt.Errorf("try get id(%v) data cache[%v] db func is nil", strID, prefixCacheForID)
	}

	data, err := LockGetOrInsertCache(s, prefixCacheForID, strID, func() (interface{}, error) {

		dataCache, find, err := s.GetFunc(id, keys)
		if err != nil {
			return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
		}

		if !find {
			// 没找到就插入
			return dbInsertFunc()
		}

		return dataCache, nil
	})
	if err != nil {
		return "", fmt.Errorf("try get and insert id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
	}

	return data, nil
}

type RedisHandle struct {
	tx           *redis.Tx       // 中间值
	setRedisData []*SetRedisInfo // 中间值
}

type MysqlHandle struct {
	xormTx *xorm.Session
}

type Session struct {
	RedisHandle
	MysqlHandle
	cacheHandler *HandleCacheInfo
}

func (s *Session) GetSetRedisInfo() []*SetRedisInfo {
	return s.setRedisData
}

func TransactionCache(ctx context.Context, s *HandleCacheInfo, transaction func(session *Session) error) error {
	sessionData := &Session{
		cacheHandler: s,
	}

	err := LockUpdateCache(ctx, sessionData, s.Conn, transaction)
	if err != nil {
		return err
	}
	return nil
}

func TransactionDB(ctx context.Context, s *HandleCacheInfo, transport func(session *Session) error) error {

	if s == nil || s.Engine == nil {
		return fmt.Errorf("sql engine is nil")
	}

	_, err := s.Engine.Transaction(func(xormS *xorm.Session) (interface{}, error) {
		sessionData := &Session{
			cacheHandler: s,
			MysqlHandle:  MysqlHandle{xormTx: xormS},
		}

		err := transport(sessionData)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

// 只更新redis
func (s *Session) Update(key string, id string, updateFunc func(data interface{}) (interface{}, error)) error {

	data, err := s.tx.HGet(s.cacheHandler.Conn.Context(), key, id).Result()
	if err != nil {
		return err
	}

	if s.cacheHandler.Stats != nil {
		s.cacheHandler.Stats.IncrementHit() // 命中
	}
	dataUnmarshal, err := s.cacheHandler.Encrypt.Unmarshal(data)
	if err != nil {
		return err
	}

	commData, err := updateFunc(dataUnmarshal)
	if err != nil {
		return err
	}

	newData, err := s.cacheHandler.Encrypt.Marshal(commData)
	if err != nil {
		return err
	}

	s.setRedisData = append(s.setRedisData, &SetRedisInfo{
		Type: OperateType_update,
		Key:  key,
		Id:   id,
		Data: newData,
	})
	return nil
}

// 回写数据库
func (s *Session) Back(key string, id string, dbUpdateFunc func(tx *xorm.Session, data interface{}) error) error {

	ctx := s.cacheHandler.Conn.Context()
	conn := s.cacheHandler.Conn

	data, err := conn.HGet(ctx, key, id).Result()
	if err != nil {
		return err
	}

	if s.cacheHandler.Stats != nil {
		s.cacheHandler.Stats.IncrementHit() // 命中
	}
	dataUnmarshal, err := s.cacheHandler.Encrypt.Unmarshal(data)
	if err != nil {
		return err
	}

	err = dbUpdateFunc(s.xormTx, dataUnmarshal)
	if err != nil {
		return err
	}
	// 只要mysql删除成功了, redis删除失败也无所谓
	conn.HDel(ctx, key, id)

	return nil
}

// func (s *HandleCacheInfo) Update(id int64, updateData func(data interface{}) (interface{}, error), keys ...interface{}) (interface{}, error) {
// 	strID := strconv.FormatInt(id, 10)
// 	for _, key := range keys {
// 		strKey, err := judgeKeyType(key)
// 		if err != nil {
// 			return nil, err
// 		}
// 		strID += "." + strKey
// 	}

// 	prefixCacheForID := prefixCache + s.ServiceName + "." + s.CacheName

// 	if s.GetFunc == nil || s.UpdateFunc == nil {
// 		return nil, fmt.Errorf("try get id(%v) data cache[%v] db func is nil", strID, prefixCacheForID)
// 	}

// 	updateFunc := func(data interface{}) (interface{}, error) {
// 		return updateData(data)
// 	}
// 	insertFunc := func() (interface{}, error) {
// 		dataCache, err := s.GetFunc(id, keys)
// 		if err != nil {
// 			return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
// 		}
// 		return dataCache, nil
// 	}

// 	return LockUpdateOrInsertCache(s, prefixCacheForID, strID, updateFunc, insertFunc)
// }
