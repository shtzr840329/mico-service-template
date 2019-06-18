# template

## 与Kratos模板项目差异
1. [+] `api/discovery.proto`: 注册中心的proto模板（NOTE：需要`kratos tool protoc api/discovery.proto`重新生成一下）
2. [#] `cmd/main.go`: 初始化服务注册构造器（L28，参照[这里](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/warden-resolver.md#%E4%BD%BF%E7%94%A8discovery)）
2. [+] `internal/server/common.go`: 注册中心微服务获取方法`DiscoveryService`的定义位置
3. [#] `internal/server/grpc/server.go`: 添加`RegisterGRPCService`方法用于向注册中心注册该项目的微服务（demo.service）
4. [#] `internal/server/http/server.go`: 添加`RegisterHTTPService`方法用于向注册中心批量添加HTTP路由
5. [+] `internal/utils/file.go`: 定义从微服务接口的swagger说明文档JSON中读取所有HTTP路由（paths字段）的函数`PickPathsFromSwaggerJSON`（主要用于上述注册HTTP接口方法`RegisterHTTPService`）

## 启动命令
`go run cmd/main.go -conf configs/ -discovery.nodes X.X.X.X:7171`
> `X.X.X.X:7171`是Bl服务发现的地址和端口