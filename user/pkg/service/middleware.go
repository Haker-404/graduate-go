package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
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
