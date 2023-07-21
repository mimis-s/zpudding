package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mimis-s/zpudding/tools/auto_tool/global"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/cmd"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/common"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/proto"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/proto_server"
	"github.com/mimis-s/zpudding/tools/auto_tool/tpl/src/services"
)

// go mod生成
func GoModGenerate() error {

	cmd := exec.Command("/bin/sh", "global/init_sh.sh", global.BootConfigYaml.OutPath, global.BootConfigYaml.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("wire genrate is err:%v", err)
		return err
	}
	fmt.Println("go mod down")
	return nil
}

func BuildShGenerate() error {
	clientPath := global.BootConfigYaml.OutPath
	tplPath := "tpl/"
	err := global.GenerateFile(clientPath, "build_proto.sh", tplPath+"build_proto.tpl")
	if err != nil {
		return err
	}
	err = global.GenerateFile(clientPath, "Makefile", tplPath+"makefile.tpl")
	if err != nil {
		return err
	}

	cmd := exec.Command("chmod", "0777", clientPath+"/build_proto.sh")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("wire genrate is err:%v", err)
		return err
	}

	cmd = exec.Command("chmod", "0777", clientPath+"/Makefile")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("wire genrate is err:%v", err)
		return err
	}

	return err
}

func main() {
	global.ParseBootYaml("global/boot.yaml")
	os.MkdirAll(global.BootConfigYaml.OutPath, 0777) // 主目录
	err := cmd.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = common.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = proto.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = proto_server.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = services.Generate()
	if err != nil {
		fmt.Println(err)
		return
	}

	// build proto
	err = BuildShGenerate()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = GoModGenerate()
	if err != nil {
		fmt.Println(err)
		return
	}
}
