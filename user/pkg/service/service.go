package service

import (
	"awesomeProject/user/pkg/model"
	"context"
)

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserList(ctx context.Context, userList []model.User) (message string, err error)
	Login(ctx context.Context, username, pwd string) (message model.Resp, err error)
}

type basicUserService struct{}

func (b *basicUserService) GetUserList(ctx context.Context, userList []model.User) (message string, err error) {
	// TODO implement the business logic of GetUserList
	return message, err
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicUserService) Login(ctx context.Context, username string, pwd string) (message model.Resp, err error) {
	// TODO implement the business logic of Login
	return message, err
}
