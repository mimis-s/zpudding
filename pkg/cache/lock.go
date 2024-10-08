package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type BagItem struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	RoleId     int64     `xorm:"not null unique(role_id_item_id) BIGINT"`
	ItemId     int       `xorm:"not null unique(role_id_item_id) INT"`
	ItemNum    int64     `xorm:"not null default 0 BIGINT"`
	ItemUseNum int64     `xorm:"not null default 0 comment('物品使用次数') BIGINT"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
	DeletedAt  time.Time `xorm:"deleted"`
}

func tryTxCache(conn *redis.Client, txf func(tx *redis.Tx) error, key, id string) error {
	const maxRetries = 1000
	for i := 0; i < maxRetries; i++ {
		err := conn.Watch(conn.Context(), txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// 乐观锁失败
			continue
		}
		return err
	}

	return fmt.Errorf("increment key[%v] id[%v] reached maximum number of retries", key, id)
}

func LockUpdateOrInsertCache(handle *HandleCacheInfo, key string, id string,
	updateFunc func(data interface{}) (interface{}, error),
	insertFunc func() (interface{}, error)) (interface{}, error) {
	// 事务函数
	var reData interface{}
	txf := func(tx *redis.Tx) error {
		var newData string
		data, err := tx.HGet(handle.Conn.Context(), key, id).Result()

		if err != nil && err != redis.Nil {
			return err
		}
		if err == redis.Nil {
			if handle.Stats != nil {
				handle.Stats.IncrementMiss() // 未命中
			}

			dataUnmarshal, err := insertFunc()
			if err != nil {
				return err
			}
			commData, err := updateFunc(dataUnmarshal)
			if err != nil {
				return err
			}
			reData = commData
			newDatav1, err := handle.Encrypt.Marshal(commData)
			if err != nil {
				return err
			}
			newData = newDatav1
		} else {
			if handle.Stats != nil {
				handle.Stats.IncrementHit() // 命中
			}
			dataUnmarshal, err := handle.Encrypt.Unmarshal(data)
			if err != nil {
				return err
			}

			commData, err := updateFunc(dataUnmarshal)
			if err != nil {
				return err
			}
			reData = commData
			newDatav1, err := handle.Encrypt.Marshal(commData)
			if err != nil {
				return err
			}
			newData = newDatav1
		}

		_, err = tx.TxPipelined(handle.Conn.Context(), func(pipe redis.Pipeliner) error {
			pipe.HSet(handle.Conn.Context(), key, id, newData)
			return nil
		})

		return err
	}

	err := tryTxCache(handle.Conn, txf, key, id)
	if err != nil {
		return "", err
	}
	return reData, nil
}

type OperateType int

var (
	OperateType_update OperateType = 1 // 只更新redis
	OperateType_back   OperateType = 2 // redis回写mysql(回写失败也不影响后续逻辑, 只是redis多了一块冗余的数据而以)

)

type SetRedisInfo struct {
	Type OperateType // 操作类型
	Key  string
	Id   string
	Data string
}

// 只更新redis
func LockUpdateCache(ctx context.Context, session *Session, conn *redis.Client,
	transaction func(session *Session) error) error {

	// 事务函数
	txf := func(tx *redis.Tx) error {

		err := transaction(session)
		if err != nil {
			return err
		}

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			for _, data := range session.GetSetRedisInfo() {
				if data.Type == OperateType_update {
					err := pipe.HSet(ctx, data.Key, data.Id, data.Data).Err()
					if err != nil {
						return err
					}
				} else {
					// 回写数据库需要将redis数据删除
					pipe.HDel(ctx, data.Key, data.Id, data.Data).Err()
				}

			}
			return nil
		})

		return err
	}

	lockKey := "update.redis.tx"

	const maxRetries = 1000
	for i := 0; i < maxRetries; i++ {
		err := conn.Watch(ctx, txf, lockKey)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// 乐观锁失败
			continue
		}
		return err
	}

	return fmt.Errorf("increment key[%v] reached maximum number of retries", lockKey)
}

func LockGetOrInsertCache(handle *HandleCacheInfo, key string, id string, calFunc func() (interface{}, error)) (interface{}, error) {
	// 事务函数
	var reData interface{}
	txf := func(tx *redis.Tx) error {

		data, err := tx.HGet(handle.Conn.Context(), key, id).Result()
		if err == nil {
			if handle.Stats != nil {
				handle.Stats.IncrementHit()
			}
			commData, err := handle.Encrypt.Unmarshal(data)
			if err != nil {
				return err
			}
			reData = commData
			return nil
		}

		if err != redis.Nil {
			return err
		}

		if handle.Stats != nil {
			handle.Stats.IncrementMiss()
		}

		commData, err := calFunc()
		if err != nil {
			return err
		}

		newData, err := handle.Encrypt.Marshal(commData)
		if err != nil {
			return err
		}
		reData = commData

		_, err = tx.TxPipelined(handle.Conn.Context(), func(pipe redis.Pipeliner) error {
			pipe.HSet(handle.Conn.Context(), key, id, newData)
			return nil
		})
		return err
	}

	err := tryTxCache(handle.Conn, txf, key, id)
	if err != nil {
		return "", err
	}
	return reData, nil
}

func LockGetCache(handle *HandleCacheInfo, key string, id string, calFunc func() (interface{}, bool, error)) (interface{}, bool, error) {
	// 事务函数
	var reData interface{}
	var bFind bool
	txf := func(tx *redis.Tx) error {

		data, err := tx.HGet(handle.Conn.Context(), key, id).Result()
		if err == nil {
			if handle.Stats != nil {
				handle.Stats.IncrementHit()
			}
			commData, err := handle.Encrypt.Unmarshal(data)
			if err != nil {
				return err
			}
			reData = commData
			return nil
		}

		if err != redis.Nil {
			return err
		}

		if handle.Stats != nil {
			handle.Stats.IncrementMiss()
		}

		commData, find, err := calFunc()
		if err != nil {
			return err
		}
		if !find {
			bFind = find
			return nil
		}

		newData, err := handle.Encrypt.Marshal(commData)
		if err != nil {
			return err
		}
		reData = commData

		_, err = tx.TxPipelined(handle.Conn.Context(), func(pipe redis.Pipeliner) error {
			pipe.HSet(handle.Conn.Context(), key, id, newData)
			return nil
		})
		return err
	}

	err := tryTxCache(handle.Conn, txf, key, id)
	if err != nil {
		return "", bFind, err
	}
	return reData, bFind, nil
}

func LockDelCache(handle *HandleCacheInfo, key string, id string, calFunc func(interface{}) error) error {
	// 事务函数
	txf := func(tx *redis.Tx) error {

		data, err := tx.HGet(handle.Conn.Context(), key, id).Result()
		if err == redis.Nil {
			if handle.Stats != nil {
				handle.Stats.IncrementMiss()
			}
			return nil
		}
		if err != nil {
			return err
		}
		if handle.Stats != nil {
			handle.Stats.IncrementHit()
		}
		newData, err := handle.Encrypt.Unmarshal(data)
		if err != nil {
			return err
		}

		err = calFunc(newData)
		if err != nil {
			return err
		}

		_, err = tx.TxPipelined(handle.Conn.Context(), func(pipe redis.Pipeliner) error {
			pipe.HDel(handle.Conn.Context(), key, id)
			return nil
		})
		return err
	}

	return tryTxCache(handle.Conn, txf, key, id)
}
