
### flatbuffers
* 新增一个build.sh文件, 可以直接编译flatc, 并添加到环境变量中
* flatbuffer源码, 修改了关于golang, go-grpc部分, 新增一个go-grpc文件

#### flatbuffer-golang结构改动
*将之前需要table <-> 后缀"T" <-> []byte 这三种结构互相转换, 替换成了外部只需要后缀"T"结构, 加上pkg/flac/codec中的序列化函数
```
例:
在fbs文件中定义如下结构
table HelloRequest {
  name:string;
  list_id:[int32];
}
flatc自动生成以下两种结构:
type Request struct {
	_tab flatbuffers.Table
}
type RequestT struct {
	MsgId uint16 `json:"msg_id"`
	Data []byte `json:"data"`
}
之前的转换[]byte -> RequestT:
1: 先通过公共函数GetRootAs...将[]byte转换为Request结构
2: 调用Request的UnPack函数, 转换为RequestT结构
改变之后的转换(pkg/flac/codec中的序列化函数):
1: 在源码中增加RequestT的成员函数MarshalTable(flatc编译时自动完成)
2: codec.Unmarshal()将[]byte转换为RequestT

```
#### flatbuffer-grpc结构改动
* 源码中的grpc调用复杂, 外部调用需要写很多序列化和反序列化代码:
```
定义一个grcp调用, 则生成如下代码:
OnMessage(ctx context.Context, in *flatbuffers.Builder,opts ...grpc.CallOption) (*Response, error)
这样的函数调用迫使外部调用者需要将结构转化为flatbuffers.Builder, 而且返回给用户的还是Response这种结构, 复杂且不实用
```
对于grcp做出如下修改:
```
OnMessage(ctx context.Context, in *RequestT,opts ...grpc.CallOption) (*ResponseT, error)
将输入输出都统一改成了我们熟悉的golang结构体, 外部调用简单直接, 不需要之前复杂冗余的解析了
```
* 新增一个grcp文件， 主要是方便外部调用做的封装, 尽可能让这些操作变成自动化
