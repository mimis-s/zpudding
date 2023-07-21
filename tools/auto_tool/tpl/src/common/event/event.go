package event

import (
	"github.com/mimis-s/zpudding/tools/auto_tool/global"
)

// 生成模板
func Generate() error {
	dbPath := global.BootConfigYaml.OutPath + "/src/common/event"
	tplPath := "tpl/src/common/event/"
	err := global.GenerateFile(dbPath, "event_define.go", tplPath+"event_define.tpl")
	if err != nil {
		return err
	}

	err = global.GenerateFile(dbPath, "event_struct.go", tplPath+"event_struct.tpl")
	if err != nil {
		return err
	}

	err = global.GenerateFile(dbPath, "event.go", tplPath+"event.tpl")
	if err != nil {
		return err
	}

	return err
}
