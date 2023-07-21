package proto

import "github.com/mimis-s/zpudding/tools/auto_tool/global"

var protoTpl = `
syntax = "proto3";

package im_home_proto;

option go_package = "{{.Name}}/src/common/commonproto/im_home_proto";

`

// 生成模板
func Generate() error {
	protoPath := global.BootConfigYaml.OutPath + "/src/proto"
	protoTplPath := "tpl/src/proto/"

	err := global.GenerateFile(protoPath, "errors.proto", protoTplPath+"errors.tpl")
	if err != nil {
		return err
	}

	// 生成对应服务的proto文件
	for _, service := range global.BootConfigYaml.Services {
		err := global.GenerateText(protoPath, "home_"+service.Tag+".proto", protoTpl)
		if err != nil {
			return err
		}
	}
	return nil
}
