package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type cacheBase struct {
	Conn      *redis.Client
	CacheLock bool       // 操作缓存的时候加分布式锁
	Stats     *CacheStat // 缓存命中率统计
	Encrypt   Serialize  // 编解码方式(如果客户端不传则为string)
	Lifetime  int        // 存活时间(s)
}

func NewBase(conn *redis.Client, stats *CacheStat, encrypt Serialize, cacheLock bool) *cacheBase {
	return &cacheBase{
		Conn:      conn,
		Stats:     stats,
		Encrypt:   encrypt,
		CacheLock: cacheLock,
	}
}

func (c *cacheBase) update(key string, id string,
	updateFunc func(data interface{}) (interface{}, error),
	getFunc func() (interface{}, error)) (interface{}, error) {
	redisKey := key + "." + id
	var reData interface{}
	data, err := c.Conn.Get(c.Conn.Context(), redisKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		if c.Stats != nil {
			c.Stats.IncrementMiss() // 未命中
		}

		dataUnmarshal, err := getFunc()
		if err != nil {
			return nil, err
		}

		commData, err := updateFunc(dataUnmarshal)
		if err != nil {
			return nil, err
		}

		reData = commData
		newDatav1, err := c.Encrypt.Marshal(commData)
		if err != nil {
			return nil, err
		}

		reData = newDatav1
	} else {
		if c.Stats != nil {
			c.Stats.IncrementHit() // 命中
		}

		dataUnmarshal, err := c.Encrypt.Unmarshal(data)
		if err != nil {
			return nil, err
		}

		commData, err := updateFunc(dataUnmarshal)
		if err != nil {
			return nil, err
		}

		reData = commData
		newDatav1, err := c.Encrypt.Marshal(commData)
		if err != nil {
			return nil, err
		}
		reData = newDatav1
	}
	_, err = c.Conn.Set(c.Conn.Context(), redisKey, reData, time.Second*time.Duration(c.Lifetime)).Result()
	return reData, err
}

func (c *cacheBase) get(key string, id string, calFunc func() (interface{}, bool, error)) (interface{}, bool, error) {
	redisKey := key + "." + id
	var reData interface{}
	data, err := c.Conn.Get(c.Conn.Context(), redisKey).Result()
	if err == nil {
		if c.Stats != nil {
			c.Stats.IncrementHit()
		}
		commData, err := c.Encrypt.Unmarshal(data)
		if err != nil {
			return nil, false, err
		}
		reData = commData
		return reData, true, nil
	}
	if err != redis.Nil {
		return nil, false, err
	}
	if c.Stats != nil {
		c.Stats.IncrementMiss()
	}
	commData, find, err := calFunc()
	if err != nil {
		return nil, false, err
	}
	if !find {
		return nil, false, nil
	}
	newData, err := c.Encrypt.Marshal(commData)
	if err != nil {
		return nil, false, err
	}
	reData = commData
	_, err = c.Conn.Set(c.Conn.Context(), redisKey, newData, time.Second*time.Duration(c.Lifetime)).Result()
	return reData, find, err
}

func (c *cacheBase) del(key string, id string, calFunc func(interface{}) error) error {
	redisKey := key + "." + id
	data, err := c.Conn.Get(c.Conn.Context(), redisKey).Result()
	if err == redis.Nil {
		if c.Stats != nil {
			c.Stats.IncrementMiss()
		}
		return nil
	}
	if err != nil {
		return err
	}
	if c.Stats != nil {
		c.Stats.IncrementHit()
	}
	newData, err := c.Encrypt.Unmarshal(data)
	if err != nil {
		return err
	}
	err = calFunc(newData)
	if err != nil {
		return err
	}
	c.Conn.Del(c.Conn.Context(), redisKey)
	return nil

}

func (c *cacheBase) UpdateCache(key string, id string,
	updateFunc func(data interface{}) (interface{}, error),
	getFunc func() (interface{}, error)) (interface{}, error) {
	if c.CacheLock {
		lockKey := fmt.Sprintf("%v.%v", key, id)
		err := nonPreemptiveLock(c.Conn, lockKey)
		if err != nil {
			return nil, err
		}
		defer nonPreemptiveUnLock(c.Conn, lockKey)
		return c.update(key, id, updateFunc, getFunc)
	}
	return c.update(key, id, updateFunc, getFunc)
}

func (c *cacheBase) InsertCache(key string, id string, data interface{}) error {
	redisKey := key + "." + id
	if c.CacheLock {
		lockKey := fmt.Sprintf("%v.%v", key, id)
		err := nonPreemptiveLock(c.Conn, lockKey)
		if err != nil {
			return err
		}
		defer nonPreemptiveUnLock(c.Conn, lockKey)
		_, err = c.Conn.Set(c.Conn.Context(), redisKey, data, time.Second*time.Duration(c.Lifetime)).Result()
		return err
	}
	newDataStr, err := c.Encrypt.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.Conn.Set(c.Conn.Context(), redisKey, newDataStr, time.Second*time.Duration(c.Lifetime)).Result()
	return err
}

func (c *cacheBase) GetCache(key string, id string, calFunc func() (interface{}, bool, error)) (interface{}, bool, error) {
	if c.CacheLock {
		lockKey := fmt.Sprintf("%v.%v", key, id)
		err := nonPreemptiveLock(c.Conn, lockKey)
		if err != nil {
			return nil, false, err
		}
		defer nonPreemptiveUnLock(c.Conn, lockKey)
		return c.get(key, id, calFunc)
	}
	return c.get(key, id, calFunc)
}

func (c *cacheBase) DelCache(key string, id string, calFunc func(interface{}) error) error {
	if c.CacheLock {
		lockKey := fmt.Sprintf("%v.%v", key, id)
		err := nonPreemptiveLock(c.Conn, lockKey)
		if err != nil {
			return err
		}
		defer nonPreemptiveUnLock(c.Conn, lockKey)
		return c.del(key, id, calFunc)
	}
	return c.del(key, id, calFunc)
}
