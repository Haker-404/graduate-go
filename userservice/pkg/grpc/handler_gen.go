// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "awesomeProject/userservice/pkg/endpoint"
	pb "awesomeProject/userservice/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getUserList grpc.Handler
	login       grpc.Handler
	getUserInfo grpc.Handler
	sign        grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.UserServiceServer {
	return &grpcServer{
		getUserList: makeGetUserListHandler(endpoints, options["GetUserList"]),
		login:       makeLoginHandler(endpoints, options["Login"]),
		getUserInfo: makeGetUserInfoHandler(endpoints, options["GetUserInfo"]),
		sign:        makeSignHandler(endpoints, options["Sign"]),
	}
}
