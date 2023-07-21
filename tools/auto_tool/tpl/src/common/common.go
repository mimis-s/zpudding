package common

import (
	"fmt"
	"os"

	"github.com/mimis-s/zpudding/tools/auto_tool/global"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common/boot_config"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common/common_client"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common/dbmodel"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common/event"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common/registry"
)

func Generate() error {
	err := boot_config.Generate() // boot_config
	if err != nil {
		return nil
	}

	err = common_client.Generate() // common_client
	if err != nil {
		return nil
	}

	err = dbmodel.Generate() // dbmodel
	if err != nil {
		fmt.Printf("generated: %v/ is err:%v\n", "dbmodle", err)
		return err
	}

	err = event.Generate() // event
	if err != nil {
		fmt.Printf("generated: %v/ is err:%v\n", "event", err)
		return err
	}

	err = registry.Generate() // registry
	if err != nil {
		fmt.Printf("generated: %v/ is err:%v\n", "registry", err)
		return err
	}

	// 创建commonproto文件夹,存储生成的proto文件
	filePath := global.BootConfigYaml.OutPath + "/src/common/commonproto"
	err = os.MkdirAll(filePath, 0777)
	if err != nil {
		fmt.Printf("generated: %v/ is err:%v\n", filePath, err)
		return err
	}
	return nil
}
