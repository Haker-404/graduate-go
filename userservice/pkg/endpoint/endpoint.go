package endpoint

import (
	model "awesomeProject/userservice/pkg/model"
	service "awesomeProject/userservice/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetUserListRequest collects the request parameters for the GetUserList method.
type GetUserListRequest struct {
}

// GetUserListResponse collects the response parameters for the GetUserList method.
type GetUserListResponse struct {
	UserList []model.User `json:"userList"`
	Message  model.Resp   `json:"message"`
}

// MakeGetUserListEndpoint returns an endpoint that invokes GetUserList on the service.
func MakeGetUserListEndpoint(s service.UserserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		userList, message := s.GetUserList(ctx)
		return GetUserListResponse{
			Message:  message,
			UserList: userList,
		}, nil
	}
}

// Failed implements Failer.

// LoginRequest collects the request parameters for the Login method.
type GetUserInfoRequest struct {
	Seq string `json:"seq"`
}

// LoginResponse collects the response parameters for the Login method.
type GetUserInfoResponse struct {
	User    model.User `json:"user"`
	Message model.Resp `json:"message"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeGetUserInfoEndpoint(s service.UserserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserInfoRequest)
		user, message := s.GetUserInfo(ctx, req.Seq)
		return GetUserInfoResponse{
			User:    user,
			Message: message,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserInfoResponse) Failed() error {
	return nil
}

//sign
type SignRequest struct {
	Seq string `json:"seq"`
}

// LoginResponse collects the response parameters for the Login method.
type SignResponse struct {
	Message model.Resp `json:"message"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeSignEndpoint(s service.UserserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignRequest)
		message := s.Sign(ctx, req.Seq)
		return SignResponse{
			Message: message,
		}, nil
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Message model.Resp `json:"message"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.UserserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		message, err := s.Login(ctx, req.Username, req.Pwd)
		return LoginResponse{
			Message: message,
		}, err
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return nil
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUserList implements Service. Primarily useful in a client.
func (e Endpoints) GetUserList(ctx context.Context) (userList []model.User, message model.Resp) {
	request := GetUserListRequest{}
	response, err := e.GetUserListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserListResponse).UserList, response.(GetUserListResponse).Message
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, username string, pwd string) (message model.Resp, err error) {
	request := LoginRequest{
		Pwd:      pwd,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Message, nil
}
func (e Endpoints) GetUserInfo(ctx context.Context, seq string) (user model.User, message model.Resp) {
	request := GetUserInfoRequest{
		Seq: seq,
	}
	response, err := e.GetUserInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserInfoResponse).User, response.(GetUserInfoResponse).Message
}
