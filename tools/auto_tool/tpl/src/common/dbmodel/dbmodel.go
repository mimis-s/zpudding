package dbmodel

import (
	"github.com/mimis-s/zpudding/pkg/lib/file"
	"github.com/mimis-s/zpudding/tools/auto_tool/global"
)

// 生成模板
func Generate() error {
	dbPath := global.BootConfigYaml.OutPath + "/src/common/dbmodel"
	tplPath := "tpl/src/common/dbmodel/"
	err := global.GenerateFile(dbPath, "table_struct.go", tplPath+"table_struct.tpl")
	if err != nil {
		return err
	}

	err = global.GenerateFile(dbPath, "table_sub.go", tplPath+"table_sub.tpl")
	if err != nil {
		return err
	}

	// 复制reverse的tpl文件
	err = file.DirCopy(tplPath+"reverse", dbPath+"/reverse")
	if err != nil {
		return err
	}

	return err
}
