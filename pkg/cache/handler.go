package cache

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	prefixCache = "cache.data.int."
)

type HandleCacheInfo struct {
	ServiceName string // 服务名字
	CacheName   string // 缓存数据的名字
	Conn        *redis.Client
	Stats       *CacheStat       // 缓存命中率统计
	Encrypt     Serialize        // 编解码方式(如果客户端不传则为string)
	GetFunc     GetFuncHanel     // db的读写函数
	UpdateFunc  UpdateFuncHandle // db的读写函数
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

func (s *HandleCacheInfo) LoadToCache(id int64, datas interface{}, keys ...interface{}) (interface{}, error) {
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

	reData, err := LockGetOrInsertCache(s, prefixCacheForID, strID, func() (interface{}, error) {
		var dataCache interface{}
		var err error
		if datas != "" {
			dataCache = datas
		} else {
			dataCache, err = s.GetFunc(id, keys)
			if err != nil {
				return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
			}
		}
		return dataCache, nil
	})
	if err != nil {
		return "", fmt.Errorf("try set id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
	}

	return reData, nil
}

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

func (s *HandleCacheInfo) Get(id int64, keys ...interface{}) (interface{}, error) {
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

		dataCache, err := s.GetFunc(id, keys)
		if err != nil {
			return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
		}
		return dataCache, nil
	})
	if err != nil {
		return "", fmt.Errorf("try get and insert id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
	}

	return data, nil
}

func (s *HandleCacheInfo) Update(id int64, updateData func(data interface{}) (interface{}, error), keys ...interface{}) (interface{}, error) {
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

	updateFunc := func(data interface{}) (interface{}, error) {
		return updateData(data)
	}
	insertFunc := func() (interface{}, error) {
		dataCache, err := s.GetFunc(id, keys)
		if err != nil {
			return "", fmt.Errorf("try get id(%v) data cache[%v] is err:%v", strID, prefixCacheForID, err)
		}
		return dataCache, nil
	}

	return LockUpdateOrInsertCache(s, prefixCacheForID, strID, updateFunc, insertFunc)
}
