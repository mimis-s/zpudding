# 项目概述
zpudding以本人+我的小猫命名 主打一个宠爱, 
它是一个集成了服务器开发工具和自动化生成编译服务器底层代码的微服务框架

可以直接使用tools的自动化生成工具生成一个完整的服务器框架, 快速开发

# 工具集成

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

8.  pkg/cache是用redis给db操作做了一个二级缓存, 和其它项目缓存的区别在于, 我们这个是当玩家在线的时候保证一些重要数据一直是热数据, 由于项目中的高频地图行为开发的缓存库

9.  cmd/flatbuffers是在flatbuffers源码的基础上进行修改, 改变之后的使用接近protoc-gen-rpcx的效果(flatbuffer+grpc)

10.  pkg/flatc/codec是flatbuffer对应的序列化/反序列化工具

10.  pkg/flatc/local_grpc_stream是flatbuffer流stream的本地调用库(all_in_one环境)

11.  pkg/flatc/pools是grpc的一个客户端池子, 主要连接grpc服务器(这里的服务发现是直接连接service的域名, 通过域名解析达到和注册中心一样的效果), 所以flat+grpc这一套不需要考虑服务ip的租期续约等操作


# 自动化生成工具
路径: https://github.com/mimis-s/zpudding/tools/auto_tool
```
自动化生成项目代码,目录工具, 用户可自定义生成项目路径及名称和想要生成的哪些服务
可在yaml文件中自定义内容: https://github.com/mimis-s/zpudding/tools/auto_tool/global/boot.yaml
outpath: "/home/zhangbin/work/pro/src/go_pro/cmd_test" // 生成项目的路径
name: "cmd_test" // 项目名(用于生成go mod, 例如:github.com/mimis-s/test)
services:
  - tag: "account"
    // 下面这些项目现在是固定生成的, 不受配置影响, 后面有需求会考虑能自由添加项目
    templateIds: ["dao", "job", "service", "server", "view"] 
  - tag: "bag"
    templateIds: ["dao", "job", "service"]
```

# 优化调整
[]待优化
问题1:现在客户端一个操作的req, 需要在客户端和服务器的proto文件里面定义两遍这个req和res, 不是特别方便
优化1:将crc32生成msgid的方式改成手动设置msgid(由00 | 000两位服务id+三位消息id构成), 每个服务新增一个OnMessage函数, 里面只有一个msgid+[]byte数据, gateway服针对msgid的前两位判断哪个服, 将客户端消息直接丢过去, 到对应的服上进行解析

[]待优化
问题2: 现在单个服的逻辑,数据库这些组件其实是高度耦合的, 虽然在一个服管理数据少的情况下是合理的, 但是一旦单个微服务里面的数据,逻辑变得复杂, 则现在的微服务内部设计就不行了
优化2: 对于绝大多数服务场景来说, 这部分就没必要优化了, 只是针对大型微服务来说, 可以将微服务用DDD领域驱动继续往下分领域,实体层进行细分

[]待优化
问题3: 现在的mq就是单纯的抛出事件, 消费事件, 不涉及当事件失败之后的处理, 这样会导致由于某些原因, 当前事件出错了, 但是mq还是把消息消费了, 这个时候就无法挽回了
优化3: 考虑为mq加入死信队列, 重试队列, 还要保证消息的幂等性, 这样保证消息被正常消费


