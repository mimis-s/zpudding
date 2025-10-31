import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	cacheKeyNonPreemptiveLock   = "dcs_non_preemp"
	nonPreemptiveLockMaxSeconds = 10
	nonPreemptiveLockTryWait    = time.Millisecond * 50
	nonPreemptiveLockTryNum     = nonPreemptiveLockMaxSeconds * time.Second / nonPreemptiveLockTryWait
)

// 非抢占式锁获取
func nonPreemptiveLock(conn *redis.Client, lockKey string) error {
	realKey := fmt.Sprintf("%v.%v", cacheKeyNonPreemptiveLock, lockKey)
	var ok bool
	var err error
	for i := 0; i < int(nonPreemptiveLockTryNum); i++ {
		ok, err = conn.SetNX(conn.Context(), realKey, 1, time.Second*nonPreemptiveLockMaxSeconds).Result()
		if ok {
			return nil
		}
		time.Sleep(nonPreemptiveLockTryWait)
	}
	if err == nil {
		err = errors.New("try lock timeout")
	}
	log.Printf("xcachedcs non preemptive lock:%v fail, err:%v", lockKey, err)
	return err
}

// 非抢占式锁释放
func nonPreemptiveUnLock(conn *redis.Client, lockKey string) {
	realKey := fmt.Sprintf("%v.%v", cacheKeyNonPreemptiveLock, lockKey)
	intCmd := conn.Del(conn.Context(), realKey)
	if err := intCmd.Err(); err != nil {
		log.Printf("xcachedcs non preemptive unlock:%v fail, err:%v", lockKey, err)
	}
}
