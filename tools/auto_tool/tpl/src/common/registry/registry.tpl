package registry

import (
	"time"

	"{{$.Name}}/src/common/boot_config"
	"{{$.Name}}/src/common/common_client"
	"{{$.Name}}/src/common/event"
	{{range $_, $m := .Services}}"{{$.Name}}/src/services/{{$m.Tag}}/api_{{$m.Tag}}"
    {{end}}
	"github.com/mimis-s/zpudding/pkg/app"
	"github.com/mimis-s/zpudding/pkg/zlog"
)

func NewDefRegistry() *app.Registry {
	s := app.NewRegistry(
		app.AddRegistryBootConfigFile(boot_config.BootConfigData),
		app.AddRegistryExBootFlags(boot_config.CustomBootFlagsData),
		app.AddRegistryGlobalCmdFlag(boot_config.BootFlagsData),
	)

	s.AddInitTask("初始化rpc调用客户端", func() error {

		// 初始化rpc调用代码
		etcdAddrs := boot_config.BootConfigData.Etcd.Addrs
		timeout := time.Duration(boot_config.BootConfigData.Etcd.Timeout * int(time.Second))
		etcdBasePath := boot_config.BootConfigData.Etcd.EtcdBasePath
		isLocal := boot_config.BootConfigData.IsLocal

	    {{range $_, $m := .Services}}api_{{$m.Tag}}.SingleNew{{$m.TagTitleCase}}Client(etcdAddrs, timeout, etcdBasePath, isLocal)
        {{end}}

		// 日志
		zlog.NewLogger(boot_config.BootConfigData.Log.Path + "/" + "log.log")

		// 初始化消息队列生产者
		// url := "amqp://dev:dev123@localhost:5672/"
		// durable := false
		err := event.InitProducers(boot_config.BootConfigData.MQ.Url, boot_config.BootConfigData.MQ.Durable)
		if err != nil {
			panic(err)
		}

		// 注册数据库
		common_client.RegisterParseMysql(boot_config.BootConfigData.DB)
		return nil
	})

	return s
}