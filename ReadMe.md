这是一个毕业设计
内部采用rpc通信 同时向外部暴露Restful API接口
go kit作为微服务工具提供者 go get 
truss用以自动生成kit代码
grpc作为调用框架
go get -u github.com/golang/protobuf
自动生成grpc pb文件:
protoc --go_out=plugins=grpc:. *.proto   //"."和"*"之间有个空格，不然会出错
protoc -I=proto --go_out=plugins=grpc:. base.proto

go get github.com/kujtimiihoxha/kit :当我们构建go-kit微服务时，会发现不同的微服务之间会有大量的冗余代码，Endpoint和Transport以及其他组件代码基本一致，书写这部分代码会浪费大量的时间，并且容易出现问题。当然，我们可以提取共用部分减少代码的冗余，但并不是一件容易的事情，十分考验我们的代码能力。那么，有没有一个工具可以快速帮助我们生成这些共用代码从而使我们集中精力关注业务逻辑呢？                                 
                                     GoKit-CLI就是解决这个问题的
Grpc-gateway
是什么
grpc-gateway is a plugin of protoc. It reads gRPC service definition, and generates a reverse-proxy server which translates a RESTful JSON API into gRPC. This server is generated according to custom options in your gRPC definition.
grpc-gateway是protoc的一个插件。它读取gRPC服务定义，并生成一个反向代理服务器，将RESTful JSON API转换为gRPC。此服务器是根据gRPC定义中的自定义选项生成的。

安装
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
protoc -I. -I%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis --grpc-gateway_out=logtostderr=true:. gateway.proto
protoc -I. -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \ --plugin=protoc-gen-go=%GOPATH%/bin/protoc-gen-go \ --grpc-gateway_out=logtostderr=true:. user.proto
用法：kit new service user 创建服务
在pkg/service 中定义业务函数 model中定义实体类
kit g s todo -w --gorilla 生成服务
-w 生成一些默认的服务中间件
--gorilla使用gorilla/mux来替换默认的http handler


很多包被墙了下不下来 使用set GOPROXY=https://goproxy.io 命令添加代理 即可下载

gopath:D:\Golang\workplace

go-kit提供以下功能：

1、Circuit breaker（熔断器）

2、Rate limiter（限流器）

3、Logging（日志）

4、Metrics（Prometheus统计）

5、Request tracing（请求跟踪）

6、Service discovery and load balancing（服务发现和负载均衡）
但随着随着业务的发展，请求量越来越大，同时，有些请求的链路也变长了，为了继续保证接口的高并发和低时延特性，团队有少量业务开始尝试GRPC。根据测试，压测一个空接口，GRPC的性能大约是HTTP＋JSON的2～3倍，在这里推荐一个压测框架fperf 。
kit g s hello -t grpc

Zipkin提供分布式追踪