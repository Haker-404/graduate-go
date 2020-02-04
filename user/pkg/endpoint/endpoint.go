package endpoint

import (
	model "awesomeProject/user/pkg/model"
	service "awesomeProject/user/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

type GetUserListRequest struct {
	UserList []model.User `json:"user_list"`
}

type GetUserListResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func MakeGetUserListEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserListRequest)
		message, err := s.GetUserList(ctx, req.UserList)
		return GetUserListResponse{
			Err:     err,
			Message: message,
		}, nil
	}
}

func (r GetUserListResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) GetUserList(ctx context.Context, userList []model.User) (message string, err error) {
	request := GetUserListRequest{UserList: userList}
	response, err := e.GetUserListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserListResponse).Message, response.(GetUserListResponse).Err
}

type LoginRequest struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

type LoginResponse struct {
	Message model.Resp `json:"message"`
	Err     error      `json:"err"`
}

func MakeLoginEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		message, err := s.Login(ctx, req.Username, req.Pwd)
		return LoginResponse{
			Err:     err,
			Message: message,
		}, nil
	}
}

func (r LoginResponse) Failed() error {
	return r.Err
}

func (e Endpoints) Login(ctx context.Context, username string, pwd string) (message model.Resp, err error) {
	request := LoginRequest{
		Pwd:      pwd,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Message, response.(LoginResponse).Err
}
