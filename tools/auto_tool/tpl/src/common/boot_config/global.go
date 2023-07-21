package boot_config

import "github.com/mimis-s/zpudding/tools/auto_tool/global"

// 生成模板
func Generate() error {
	globalPath := global.BootConfigYaml.OutPath + "/src/common/boot_config"
	tplPath := "tpl/src/common/boot_config/"
	global.GenerateFile(globalPath, "boot_config.go", tplPath+"boot_config.tpl")
	global.GenerateFile(globalPath, "init.go", tplPath+"init.tpl")

	return nil
}
