package services

var api_proto_tpl = `
syntax = "proto3";

package api_{{.ServiceName}};

service {{.TagTitleCase}} {
}
`
var dao_tpl = `
package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"{{.ProName}}/src/common/common_client"
	_"github.com/mimis-s/zpudding/pkg/dfs"
	"github.com/mimis-s/zpudding/pkg/zlog"
	"xorm.io/xorm"
)

var ProviderSet = wire.NewSet(New)

type Dao struct {
	Db         *xorm.Engine
	cache      *common_client.RedisClient
	// dfsHandler dfs.DFSHandler
}

func New() (*Dao, error) {

	// 初始化数据库xorm

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_{{.TagTitleCase}})
	if err != nil {
		zlog.Warn("{{.ServiceName}} dao new engine is err:%v", err)
		return nil, fmt.Errorf("{{.ServiceName}} dao new engine is err:%v", err)
	}

	// dfsHandler, err := dfs.NewDFSHandler(&boot_config.BootConfigData.DFS)
	// if err != nil {
	// 	panic(err)
	// }

	// // 新建桶
	// err = dfsHandler.TryMakeBucket()
	// if err != nil {
	// 	panic(err)
	// }

	dao := &Dao{
		Db:         engine,
		cache:      common_client.NewRedisClient(),
		// dfsHandler: dfsHandler,
	}

	return dao, nil
}

`
var service_tpl = `
package service

import (
	"github.com/google/wire"
	"{{.ProName}}/src/services/{{.ServiceName}}/api_{{.ServiceName}}"
	"{{.ProName}}/src/services/{{.ServiceName}}/dao"
	rpcxService "github.com/mimis-s/zpudding/pkg/rpcx/service"
)

var ProviderSet = wire.NewSet(dao.ProviderSet, NewServiceHandler)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func NewServiceHandler(rpcSvc *rpcxService.ServerManage, d *dao.Dao) (*Service, error) {

	S = &Service{
		Dao: d,
	}

	// 绑定rpcx服务
	err := api_{{.ServiceName}}.Register{{.TagTitleCase}}Service(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	return S, nil
}
`
var job_tpl = `
package job

import (
	"{{.ProName}}/src/common/boot_config"
	"{{.ProName}}/src/common/event"
	"{{.ProName}}/src/services/{{.ServiceName}}/service"
	"github.com/mimis-s/zpudding/pkg/mq/rabbitmq"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service) *Job {
	url := boot_config.BootConfigData.MQ.Url
	durable := boot_config.BootConfigData.MQ.Durable

	j := &Job{s}

	event.InitConsumers(url, durable,
		[]*rabbitmq.ConsumersQueue{
			// {event.Event_UserLogin, j.userLoginOk},
		},
	)

	return j
}
// 例子:
// func (j *Job) userLoginOk(payload interface{}) error {
// 	userLogin := payload.(*event.UserLogin)
// 	return nil
// }
`
