package services

import (
	"fmt"
	"os/exec"

	"github.com/mimis-s/zpudding/tools/auto_tool/global"
)

type ServiceInfo struct {
	ProName      string
	ServiceName  string
	TagTitleCase string
}

// 生成模板
func Generate() error {
	protoPath := global.BootConfigYaml.OutPath + "/src/services"
	tplPath := "tpl/src/services"

	for _, service := range global.BootConfigYaml.Services {
		serviceProtoPath := protoPath + "/" + service.Tag
		info := &ServiceInfo{
			ProName:      global.BootConfigYaml.Name,
			ServiceName:  service.Tag,
			TagTitleCase: service.TagTitleCase,
		}
		// proto
		err := global.GenerateTextForData(serviceProtoPath+"/api_"+service.Tag, "api_"+service.Tag+".proto", api_proto_tpl, info)
		if err != nil {
			return err
		}
		// dao
		err = global.GenerateTextForData(serviceProtoPath+"/dao", "dao.go", dao_tpl, info)
		if err != nil {
			return err
		}
		// service
		err = global.GenerateTextForData(serviceProtoPath+"/service", "service.go", service_tpl, info)
		if err != nil {
			return err
		}
		// job
		err = global.GenerateTextForData(serviceProtoPath+"/job", "job.go", job_tpl, info)
		if err != nil {
			return err
		}

		// wire
		err = global.GenerateFileForData(serviceProtoPath, "wire.go", tplPath+"/wire.tpl", info)
		if err != nil {
			return err
		}
	}

	return nil
}

// 依赖注入生成
func GenerateWire() error {
	for _, service := range global.BootConfigYaml.Services {
		serviceProtoPath := global.BootConfigYaml.OutPath + "/src/services/" + service.Tag

		// 执行wire命令
		cmd := exec.Command("wire", serviceProtoPath+"/wire.go")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("wire genrate is err:%v", err)
			return err
		}
	}
	return nil
}
