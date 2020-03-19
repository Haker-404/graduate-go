package service

import (
	model "awesomeProject/userservice/pkg/model"
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserserviceService) UserserviceService

type loggingMiddleware struct {
	logger log.Logger
	next   UserserviceService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserserviceService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserserviceService) UserserviceService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetUserList(ctx context.Context) (userList []model.User, message model.Resp) {
	defer func() {
		l.logger.Log("")
	}()
	return l.next.GetUserList(ctx)
}
func (l loggingMiddleware) Login(ctx context.Context, username string, pwd string) (message model.Resp, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "pwd", pwd, "message", message, "err", err)
	}()
	return l.next.Login(ctx, username, pwd)
}
func (l loggingMiddleware) GetUserInfo(ctx context.Context, seq string) (user model.User, message model.Resp) {

	return l.next.GetUserInfo(ctx, seq)
}
func (l loggingMiddleware) Sign(ctx context.Context, seq string) (message model.Resp) {

	return l.next.Sign(ctx, seq)
}
