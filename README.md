#### 项目概述
zpudding以本人+我的小猫命名 主打一个宠爱, 
它是一个集成了服务器开发工具和自动化生成编译服务器底层代码的微服务框架

可以直接使用tools的自动化生成工具生成一个完整的服务器框架, 快速开发

#### 工具集成

1.  cmd/protoc-gen-rpcx: 将rpcx集成到proto生成里面

2.  pkg/net网络库对外有两个接口:InitService(),Listen()
    InitService接口的传入参数是addr，协议(现在仅支持TCP)类型，消息回调函数
    Listen接口主要是开启网络服务，监听，收发消息

3.  pkg/rpcx库主要封装了rpcx服务器的创建，运行接口，客户端创建，调用接口，代码简单明了，
    详细使用可参考protobuf插件集成rpcx: https://github.com/mimis-s/zpudding/cmd/protoc-gen-rpcx

4.  pkg/dfs封装了s3,cloud storage, minio三个分布式文件存储系统的集成api

5.  pkg/rabbitmq使用topic主题交换机, 支持同时给多个消费者, 有选择性的分发消息

6.  pkg/app库是微服务框架底层服务启动框架

7.  pkg/zlog是基于lumberjack日志库二次开发的日志库, 参考链接:https://github.com/natefinch/lumberjack

#### 自动化生成
自动化生成项目代码,目录工具
