package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/mimis-s/zpudding/pkg/lib/file"
	"github.com/mimis-s/zpudding/tools/auto_tool/global"
)

// 生成模板
func Generate() error {
	cmdPath := global.BootConfigYaml.OutPath + "/cmd/all_in_one"
	mainPath := cmdPath + "/main.go"
	os.MkdirAll(cmdPath, 0777)
	mainFile, err := os.Create(mainPath)
	if err != nil {
		fmt.Println(err)
	}

	mainTemp, err := template.ParseFiles("tpl/cmd/main.tpl")
	if err != nil {
		return err
	}
	var buf bytes.Buffer

	mainTemp.Execute(&buf, global.BootConfigYaml)

	_, err = mainFile.Write(buf.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("successfully generated: cmd/all_in_one/main.go")

	// 复制配置文件
	err = file.FileCopy("tpl/cmd/boot_config.yaml", cmdPath+"/boot_config.yaml")
	if err != nil {
		return err
	}
	return nil
}
