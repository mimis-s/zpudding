package event

import (
	_ "github.com/mimis-s/zpudding/pkg/mq/rabbitmq"
)

// 示例:
// var (
// 	Event_UserLogin        = &rabbitmq.EventStruct{"user.login.ok", UserLogin{}}
// )
