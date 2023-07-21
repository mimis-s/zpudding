package common_client

import (
	"github.com/mimis-s/zpudding/tools/auto_tool/global"
)

// 生成模板
func Generate() error {
	clientPath := global.BootConfigYaml.OutPath + "/src/common/common_client"
	tplPath := "tpl/src/common/common_client/"
	global.GenerateFile(clientPath, "dfs.go", tplPath+"dfs.tpl")
	global.GenerateFile(clientPath, "mysql.go", tplPath+"mysql.tpl")
	global.GenerateFile(clientPath, "redis.go", tplPath+"redis.tpl")

	return nil
}
