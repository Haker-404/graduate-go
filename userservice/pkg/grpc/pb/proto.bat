#自动生成 pb，gw文件
protoc -I. -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. userservice.proto
protoc -I. -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. userservice.proto
protoc -I. -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. userservice.proto
