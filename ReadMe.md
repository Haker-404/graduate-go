这是一个毕业设计
go kit作为微服务工具提供者 go get 
truss用以自动生成kit代码
grpc作为调用框架
go get -u github.com/golang/protobuf
自动生成grpc pb文件:
protoc --go_out=plugins=grpc:. *.proto   //"."和"*"之间有个空格，不然会出错
protoc -I=proto --go_out=plugins=grpc:. base.proto

go get github.com/kujtimiihoxha/kit :当我们构建go-kit微服务时，会发现不同的微服务之间会有大量的冗余代码，Endpoint和Transport以及其他组件代码基本一致，书写这部分代码会浪费大量的时间，并且容易出现问题。当然，我们可以提取共用部分减少代码的冗余，但并不是一件容易的事情，十分考验我们的代码能力。那么，有没有一个工具可以快速帮助我们生成这些共用代码从而使我们集中精力关注业务逻辑呢？                                 
                                     GoKit-CLI就是解决这个问题的

很多包被墙了下不下来 使用set GOPROXY=https://goproxy.io 命令添加代理 即可下载

gopath:D:\Golang\workplace