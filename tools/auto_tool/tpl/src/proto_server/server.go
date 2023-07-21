package proto_server

import "github.com/mimis-s/zpudding/tools/auto_tool/global"

var server_protoTpl = `
// 服务器自用公共结构，客户端不要编译！
syntax = "proto3";

package im_proto_server;

option go_package = "{{.Name}}/src/common/commonproto/im_proto_server";


`

// 生成模板
func Generate() error {
	protoPath := global.BootConfigYaml.OutPath + "/src/proto_server"

	err := global.GenerateText(protoPath, "server.proto", server_protoTpl)
	if err != nil {
		return err
	}
	return nil
}
