package grpc

import (
	endpoint "awesomeProject/userservice/pkg/endpoint"
	pb "awesomeProject/userservice/pkg/grpc/pb"
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetUserListHandler creates the handler logic
func makeGetUserListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserListEndpoint, decodeGetUserListRequest, encodeGetUserListResponse, options...)
}

// decodeGetUserListResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserList request.
// TODO implement the decoder
func decodeGetUserListRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Userservice' Decoder is not impelemented")
}

// encodeGetUserListResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserListResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Userservice' Encoder is not impelemented")
}
func (g *grpcServer) GetUserList(ctx context1.Context, req *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
	_, rep, err := g.getUserList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserListReply), nil
}

// makeLoginHandler creates the handler logic
func makeLoginHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.LoginEndpoint, decodeLoginRequest, encodeLoginResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Login request.
// TODO implement the decoder
func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := endpoint.LoginRequest{
		Username: r.(*pb.LoginRequest).Username,
		Pwd:      r.(*pb.LoginRequest).Pwd,
	}

	return req, nil
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := &pb.Resp{
		Msg:  r.(endpoint.LoginResponse).Message.Msg,
		Succ: r.(endpoint.LoginResponse).Message.Succ,
	}
	rep := pb.LoginReply{
		Resp: resp,
	}
	return &rep, nil
}
func (g *grpcServer) Login(ctx context1.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginReply), nil
}

func makeGetUserInfoHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserInfoEndpoint, decodeGetUserInfoRequest, encodeGetUserInfoResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Login request.
// TODO implement the decoder
func decodeGetUserInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := endpoint.GetUserInfoRequest{
		Seq: r.(*pb.GetUserInfoRequest).Seq,
	}

	return req, nil
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := &pb.Resp{
		Msg:  r.(endpoint.GetUserInfoResponse).Message.Msg,
		Succ: r.(endpoint.GetUserInfoResponse).Message.Succ,
	}
	user := &pb.User{
		Name:  r.(endpoint.GetUserInfoResponse).User.Name,
		Seq:   r.(endpoint.GetUserInfoResponse).User.Seq,
		Photo: r.(endpoint.GetUserInfoResponse).User.Photo,
	}
	rep := pb.GetUserInfoReply{
		User: user,
		Resp: resp,
	}
	return &rep, nil
}
func (g *grpcServer) GetUserInfo(ctx context1.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoReply, error) {
	_, rep, err := g.getUserInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserInfoReply), nil
}

//sign
func makeSignHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SignEndpoint, decodeSignRequest, encodeSignResponse, options...)
}

// decodeLoginResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Login request.
// TODO implement the decoder
func decodeSignRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := endpoint.SignRequest{
		Seq: r.(*pb.SignRequest).Seq,
	}

	return req, nil
}

// encodeLoginResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSignResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := &pb.Resp{
		Msg:  r.(endpoint.SignResponse).Message.Msg,
		Succ: r.(endpoint.SignResponse).Message.Succ,
	}
	rep := pb.SignReply{
		Resp: resp,
	}
	return &rep, nil
}
func (g *grpcServer) Sign(ctx context1.Context, req *pb.SignRequest) (*pb.SignReply, error) {
	_, rep, err := g.sign.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SignReply), nil
}
