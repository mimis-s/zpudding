package registry

import (
	"github.com/google/wire"
	"{{.Name}}/src/common/boot_config"
	"github.com/mimis-s/zpudding/pkg/rpcx/service"
)

var (
	DefaultAppRpcSet = wire.NewSet(NewRpcService)
)

func NewRpcService() (*service.ServerManage, error) {

	if boot_config.BootConfigData.IsLocal {
		return nil, nil
	}

	listenAddr := boot_config.CustomBootFlagsData.RpcListenPort
	exposeAddr := boot_config.CustomBootFlagsData.RpcExposePort
	etcdAddrs := boot_config.BootConfigData.Etcd.Addrs
	etcdBasePath := boot_config.BootConfigData.Etcd.EtcdBasePath

	s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath, listenAddr)
	if err != nil {
		return nil, err
	}
	return s, nil
}
