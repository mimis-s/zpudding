package cache

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var retryGet = 3 // 当插入失败之后再重新读取的次数(一行数据的多列会被缓存, 防止并发insert失败)

// dbData都是表结构
type CacheManger interface {
	CacheGet(rid string, tableName string, dbData interface{}, cacheCol string, keys ...interface{}) (interface{}, bool, error)
	CacheUpdate(rid string, tableName string, dbData interface{}, cacheCol string, keys ...interface{}) error
	CacheInsert(rid string, tableName string, dbData interface{}) error
}

type HandleCacheInfo struct {
	tableName string       // 数据表的名字(key的前缀)
	colName   string       // 数据列的名字(缓存数据的名字)
	tableType reflect.Type // 要缓存的表
	cache     *cacheBase
	manger    CacheManger // 数据库读写管理
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

func (s *HandleCacheInfo) get(rid string, keys ...interface{}) (interface{}, bool, error) {
	colName := ""
	if s.colName != s.tableName {
		colName = s.colName
	}
	dataValue := reflect.New(s.tableType).Interface()
	dataCache, bFind, err := s.manger.CacheGet(rid, s.tableName, dataValue, colName, keys...)
	if err != nil {
		return dataCache, bFind, fmt.Errorf("try get id(%v) data cache[%v] is err:%v", s.tableName, colName, err)
	}
	if colName == "" {
		return dataCache, bFind, nil
	}
	if !bFind {
		return nil, false, nil
	}
	typeData := reflect.TypeOf(dataCache).Elem()
	valueData := reflect.ValueOf(dataCache).Elem()
	for i := 0; i < typeData.NumField(); i++ {
		field := typeData.Field(i)
		xormTags := strings.Split(field.Tag.Get("xcache"), " ")
		if len(xormTags) > 0 && xormTags[0] == colName {
			return valueData.Field(i).Interface(), bFind, nil
		}
	}
	return dataCache, bFind, fmt.Errorf("try get id(%v) data cache col:%v is not found", s.tableName, colName)
}

func (s *HandleCacheInfo) update(rid string, colData interface{}, keys ...interface{}) error {
	colName := ""
	if s.colName != s.tableName {
		colName = s.colName
	}
	dataValue := reflect.New(s.tableType).Interface()
	if colName != "" {
		typeData := reflect.TypeOf(dataValue).Elem()
		valueData := reflect.ValueOf(dataValue).Elem()
		bFindCol := false
		for i := 0; i < typeData.NumField(); i++ {
			field := typeData.Field(i)
			xormTags := strings.Split(field.Tag.Get("xcache"), " ")
			if len(xormTags) > 0 && xormTags[0] == colName {
				if valueData.Field(i).IsValid() && valueData.Field(i).CanSet() {
					valueData.Field(i).Set(reflect.ValueOf(colData))
					bFindCol = true
					break
				}
			}
		}
		if !bFindCol {
			return fmt.Errorf("try get id(%v) data cache col:%v is not found", s.tableName, colName)
		}
	}
	err := s.manger.CacheUpdate(rid, s.tableName, dataValue, colName, keys...)
	if err != nil {
		return fmt.Errorf("try get and update id(%v) data[%v] cache[%v] to db is err:%v", s.tableName, colData, colName, err)
	}
	return nil
}

func (s *HandleCacheInfo) insert(rid string, dbData interface{}, keys ...interface{}) (interface{}, error) {
	colName := ""
	if s.colName != s.tableName {
		colName = s.colName
	}
	err := s.manger.CacheInsert(rid, s.tableName, dbData)
	if err != nil {
		return nil, fmt.Errorf("try get and insert id(%v) data[%v] cache to db is err:%v", s.tableName, dbData, err)
	}

	if colName != "" {
		typeData := reflect.TypeOf(dbData).Elem()
		valueData := reflect.ValueOf(dbData).Elem()
		for i := 0; i < typeData.NumField(); i++ {
			field := typeData.Field(i)
			xormTags := strings.Split(field.Tag.Get("xcache"), " ")
			if len(xormTags) > 0 && xormTags[0] == colName {
				return valueData.Field(i).Interface(), nil
			}
		}
	}

	return dbData, nil
}

func (s *HandleCacheInfo) Get(rid string, keys ...interface{}) (interface{}, bool, error) {

	prefixCacheForID := s.tableName + "." + rid
	if s.manger == nil {
		return nil, false, fmt.Errorf("try get id(%v) data cache[%v] db func is nil", prefixCacheForID, s.colName)
	}

	data, find, err := s.cache.GetCache(prefixCacheForID, s.colName, func() (interface{}, bool, error) {
		return s.get(rid, keys...)
	})

	if err != nil {
		return nil, false, fmt.Errorf("try get id(%v) data cache[%v] is err:%v", prefixCacheForID, s.colName, err)
	}

	return data, find, nil
}

func (s *HandleCacheInfo) Update(rid string, updateData func(data interface{}) (interface{}, error), keys ...interface{}) (interface{}, error) {

	prefixCacheForID := s.tableName + "." + rid
	if s.manger == nil {
		return nil, fmt.Errorf("try update id(%v) data cache[%v] db func is nil", prefixCacheForID, s.colName)
	}

	updateFunc := func(data interface{}) (interface{}, error) {
		newData, err := updateData(data)
		if err != nil {
			return newData, err
		}
		if err := s.update(rid, newData, keys...); err != nil {
			return nil, err
		}
		return newData, nil
	}

	getFunc := func() (interface{}, error) {
		dataCache, find, err := s.get(rid, keys...)
		if err != nil {
			return "", fmt.Errorf("try update id(%v) data cache[%v] is err:%v", prefixCacheForID, s.colName, err)
		}
		if !find {
			return "", fmt.Errorf("try update id(%v) data cache[%v] is not found", prefixCacheForID, s.colName)
		}
		return dataCache, nil
	}
	return s.cache.UpdateCache(prefixCacheForID, s.colName, updateFunc, getFunc)
}

func (s *HandleCacheInfo) Insert(rid string, data interface{}, keys ...interface{}) error {
	prefixCacheForID := s.tableName + "." + rid
	if s.manger == nil {
		return fmt.Errorf("try insert id(%v) data cache[%v] db func is nil", prefixCacheForID, s.colName)
	}
	colData, err := s.insert(rid, data, keys...)
	if err != nil {
		return err
	}

	return s.cache.InsertCache(prefixCacheForID, s.colName, colData)
}
