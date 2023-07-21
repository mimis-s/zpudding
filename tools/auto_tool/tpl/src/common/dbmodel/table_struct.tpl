package dbmodel

import (
	"strconv"

	"{{.Name}}/src/common/boot_config"
)

// 分库
type DbSubTreasury string

const (
    {{range $_, $m := .Services}}ShardGroupTag_{{$m.TagTitleCase}} = "{{$m.Tag}}"
    {{end}}
)

// 分库名字
func ShardDatabaseName() map[string]bool {
	return map[string]bool{
        {{range $_, $m := .Services}}ShardGroupTag_{{$m.TagTitleCase}}: true,
        {{end}}
	}
}

// 分表
type DbTableInterface interface {
	SubName() string  // 表名
	SubTableNum() int // 分表数量
	BindSubTreasury() DbSubTreasury
}

func subName(name string, value int64, tableNum int) string {
	temp := value % int64(tableNum)
	if temp == 0 || !boot_config.BootConfigData.DataBaseShard {
		return name
	}
	return name + "_" + strconv.FormatInt(temp, 10)
}
