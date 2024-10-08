package cache

import (
	"fmt"
	"sync"
)

/*
	并发问题:
	1: 在load数据还没有读取到redis中的时候, 这个时候外部数据更新了, 要update
	1解决: 所有操作都应该给这条redis记录加一个乐观锁, update的时候发现这条记录已经被load锁住了, 就会等待load结束之后再操作
*/

type UpdateFuncHandle func(id int64, keys []interface{}, data interface{}) error
type GetFuncHanel func(id int64, keys []interface{}) (interface{}, bool, error)

type CacheInfo interface {
	/*
		这里LoadToCache的data使用可变参数是因为本来有GetFuncHanel函数,
		不传data也可以直接从get函数获取, 但是get函数可能是mysql操作,
		而一般load的时候mysql数据已经读取出来了,
		所以为了减少一次数据库操作可以传入一个参数值
	*/
	// LoadToCache(id int64, data interface{}, keys ...interface{}) (interface{}, error)                                  // 将信息读入到缓存
	LogoutCache(id int64, keys ...interface{}) error                                                                   // 将缓存数据读入数据库, 删除缓存
	Get(id int64, keys ...interface{}) (interface{}, bool, error)                                                      // 获取cache/mysql数据
	Update(id int64, updateData func(data interface{}) (interface{}, error), keys ...interface{}) (interface{}, error) // 更新cache数据
}

var mapCache sync.Map

// 每个数据对应一个typeid, 只要不重复就可以
func RegisterCacheInfoHandler(cType int, datas ...CacheInfo) {
	for _, data := range datas {
		mapCache.Store(cType, data)
	}
}

func GetCacheInfoHandler(cType int) CacheInfo {
	var data interface{}
	var ok bool
	if data, ok = mapCache.Load(cType); !ok {
		errInfo := fmt.Errorf("not found type[%v] cache, please register", cType)
		panic(errInfo)
	}
	return data.(CacheInfo)
}
