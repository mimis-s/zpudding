package registry

import "github.com/mimis-s/zpudding/tools/auto_tool/global"

// 生成模板
func Generate() error {
	schedulerPath := global.BootConfigYaml.OutPath + "/src/common/registry"
	tplPath := "tpl/src/common/registry/"
	global.GenerateFile(schedulerPath, "registry.go", tplPath+"registry.tpl")
	global.GenerateFile(schedulerPath, "rpc_service.go", tplPath+"rpc_service.tpl")

	return nil
}
