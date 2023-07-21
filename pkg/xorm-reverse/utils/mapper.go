package utils

import (
	"xorm.io/xorm/names"
	"xorm.io/xorm/schemas"
)

func GetMapperByName(mapname string) names.Mapper {
	switch mapname {
	case "gonic":
		return names.LintGonicMapper
	case "same":
		return names.SameMapper{}
	default:
		return names.SnakeMapper{}
	}
}

func GetColumnName(tables *schemas.Table, name string) bool {
	for _, c := range tables.ColumnsSeq() {
		// fmt.Printf("col:%v\n", c)
		if c == name {
			return true
		}
	}
	return false
}
