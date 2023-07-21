//go:build wireinject
// +build wireinject

package {{.ServiceName}}

import (
	"github.com/google/wire"
	"{{.ProName}}/src/common/registry"
	"{{.ProName}}/src/services/{{.ServiceName}}/job"
	"{{.ProName}}/src/services/{{.ServiceName}}/service"
	"github.com/mimis-s/zpudding/pkg/app"
	rpcxService "github.com/mimis-s/zpudding/pkg/rpcx/service"
)

func appInject(a *app.App) (interface{}, error) {
	panic(wire.Build(
		registry.DefaultAppRpcSet,
		service.ProviderSet,
		job.InitMQ,
		registerHandler,
	))
}

func registerHandler(sm *rpcxService.ServerManage, svcHandler *service.Service, _ *job.Job, a *app.App) (interface{}, error) {
	a.AddService("rpc", sm)
	return nil, nil
}

func createAppInfo() (*app.AppOutSideInfo, error) {
	var err error
	appInfo := app.NewAppOutSide("{{.ServiceName}}",
		func(a *app.App) error {
			_, err = appInject(a)
			return err

		},
	)

	return appInfo, err
}

func CreateAppInfo() (*app.AppOutSideInfo, error) {
	appInfo, err := createAppInfo()
	if err != nil {
		panic(err)
	}
	return appInfo, err
}

