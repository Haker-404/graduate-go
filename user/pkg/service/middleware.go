package service

import (
	model "awesomeProject/user/pkg/model"
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetUserList(ctx context.Context, userList []model.User) (message string, err error) {
	defer func() {
		l.logger.Log("method", "GetUserList", "userList", userList, "message", message, "err", err)
	}()
	return l.next.GetUserList(ctx, userList)
}

func (l loggingMiddleware) Login(ctx context.Context, username string, pwd string) (message model.Resp, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "pwd", pwd, "message", message, "err", err)
	}()
	return l.next.Login(ctx, username, pwd)
}
