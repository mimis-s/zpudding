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

var itemID = 1

func initHandler() *HandleCacheInfo {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       4,  // 默认DB 0
	})
	// back := func(percent float64) {
	// 	fmt.Printf("命中率:%.2f\n", percent)
	// }

	dataSourceName := "root:dev123@tcp(192.168.1.22:3380)/sim_zhangbin?charset=utf8mb4"

	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	getFunc := func(id int64, keys []interface{}) (interface{}, error) {
		t := keys[0].(int)
		fmt.Printf("tt:%v\n", t)
		dbData := &BagItem{}
		find, err := engine.Table("bag_item").Where(
			"id = ?", id).Get(dbData)
		if err != nil {
			return "", fmt.Errorf("err: %v", err)
		}
		if !find {
			return "", fmt.Errorf("not found")
		}
		return dbData, nil
	}

	updateFunc := func(id int64, keys []interface{}, data interface{}) error {
		// dbData := data.(*BagItem)
		t := keys[0].(int)
		fmt.Printf("tt:%v\n", t)
		_, err = engine.Table("bag_item").
			Where("id = ?", id).Update(data)
		if err != nil {
			return fmt.Errorf("err:%v", err)
		}
		return nil
	}
	// ("testService", "1", conn, getFunc, updateFunc, time.Second*1, back)
	handlerData := &HandleCacheInfo{
		ServiceName: "testService",
		CacheName:   "1",
		Conn:        conn,
		Encrypt: &JsonType{
			Val: &BagItem{},
		},
		// Stats:      NewCacheStat(time.Second*1, back),
		GetFunc:    getFunc,
		UpdateFunc: updateFunc,
	}
	// RegisterCacheInfoHandler(1, handlerData)
	return handlerData
}

var timesAdd int32 = 0
var timesDel int32 = 0

func TestIntCache(t *testing.T) {
	handlerData := initHandler()

	wg := sync.WaitGroup{}
	wg.Add(20)
	i := 0
	for {
		go func() {
			for j := 0; j < 1000; j++ {
				strData, err := handlerData.Update(int64(itemID), func(data interface{}) (interface{}, error) {
					itemNum := data.(*BagItem)
					itemNum.ItemNum += 10
					return itemNum, nil
				}, 29)
				if err != nil {
					panic(err)
				}
				atomic.AddInt32(&timesAdd, 1)

				itemNum := strData.(*BagItem)
				fmt.Printf("增加:%v %v\n", timesAdd, itemNum.ItemNum)
			}
			wg.Done()
		}()
		i++
		if i == 10 {
			break
		}
	}

	i = 0
	for {
		go func() {
			for j := 0; j < 1000; j++ {
				strData, err := handlerData.Update(int64(itemID), func(data interface{}) (interface{}, error) {
					itemNum := data.(*BagItem)
					itemNum.ItemNum -= 10
					return itemNum, nil
				}, 29)
				if err != nil {
					panic(err)
				}
				atomic.AddInt32(&timesDel, 1)

				itemNum := strData.(*BagItem)
				fmt.Printf("减少:%v %v\n", timesDel, itemNum.ItemNum)
			}
			wg.Done()
		}()
		i++
		if i == 10 {
			break
		}
	}

	wg.Wait()

	err := handlerData.LogoutCache(int64(itemID), 29)
	if err != nil {
		panic(err)
	}
	fmt.Printf("数据:%v %v\n", timesAdd, timesDel)
}
