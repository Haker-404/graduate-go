package endpoint

import (
	"awesomeProject/model"
	service "awesomeProject/user/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetUserListRequest collects the request parameters for the GetUserList method.
type GetUserListRequest struct {
	UserList []model.User `json:"user_list"`
}

// GetUserListResponse collects the response parameters for the GetUserList method.
type GetUserListResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err"`
}

// MakeGetUserListEndpoint returns an endpoint that invokes GetUserList on the service.
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

// Failed implements Failer.
func (r GetUserListResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUserList implements Service. Primarily useful in a client.
func (e Endpoints) GetUserList(ctx context.Context, userList []model.User) (message string, err error) {
	request := GetUserListRequest{UserList: userList}
	response, err := e.GetUserListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserListResponse).Message, response.(GetUserListResponse).Err
}
