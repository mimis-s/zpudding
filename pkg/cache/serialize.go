package cache

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// 序列化方式
type Serialize interface {
	Unmarshal(data string) (any, error)
	Marshal(v any) (string, error)
}

type Int64Type struct {
}

func (i *Int64Type) Unmarshal(data string) (any, error) {
	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (i *Int64Type) Marshal(v any) (string, error) {
	val := strconv.FormatInt(int64(v.(int64)), 10)
	return val, nil
}

type JsonType struct {
	Val interface{}
}

func (j *JsonType) Unmarshal(data string) (any, error) {
	reData := reflect.New(reflect.TypeOf(j.Val).Elem()).Interface()
	err := json.Unmarshal([]byte(data), reData)
	return reData, err
}

func (j *JsonType) Marshal(v any) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}
