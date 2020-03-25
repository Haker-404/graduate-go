package service

import (
	"awesomeProject/userservice/pkg/dao"
	"awesomeProject/userservice/pkg/model"
	jwt "awesomeProject/userservice/pkg/utils"
	"context"
	"fmt"
	"time"
)

// UserserviceService describes the service.
type UserserviceService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserList(ctx context.Context, date string, isSigned bool) (signInfoList []model.SignInfo, message model.Resp)
	Login(ctx context.Context, username, pwd string) (message model.Resp, err error)
	GetUserInfo(ctx context.Context, seq string) (user model.User, message model.Resp)
	Sign(ctx context.Context, seq string) (message model.Resp)
}

type basicUserserviceService struct{}

var d = &dao.UserDao{}

func (b *basicUserserviceService) GetUserList(ctx context.Context, date string, isSigned bool) (signInfoList []model.SignInfo, message model.Resp) {
	// TODO implement the business logic of GetUserList
	t_time, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	signInfoList, message = d.GetUserList(t_time, isSigned)
	return
}
func (b *basicUserserviceService) GetUserInfo(ctx context.Context, seq string) (user model.User, message model.Resp) {
	user, message = d.GetUserInfo(seq)
	return
}
func (b *basicUserserviceService) Sign(ctx context.Context, seq string) (message model.Resp) {
	user, message := b.GetUserInfo(ctx, seq)
	if user.Seq == "" {
		message.Succ = "302"
		message.Msg = "签到失败,无此用户！"
	} else {
		err := d.InsertSign(seq)
		if err != nil {
			message.Msg = err.Error()
		} else {
			message.Succ = "200"
			message.Msg = "签到成功"
		}

	}
	return
}
func (b *basicUserserviceService) Login(ctx context.Context, seq, pwd string) (message model.Resp, err error) {
	// TODO implement the business logic of Login
	user, message := d.GetUserInfo(seq, pwd)
	if user.Seq != "" {
		message.Succ = "200"
		token, err := jwt.CreateToken(user.Seq, user.Name, user.IsManager)
		message.Msg = token
		fmt.Print(token, err)
	} else {
		message.Msg = "登录失败，账号或者密码错误"
		message.Succ = "403"
	}
	return message, err
}

// NewBasicUserserviceService returns a naive, stateless implementation of UserserviceService.
func NewBasicUserserviceService() UserserviceService {
	return &basicUserserviceService{}
}

// New returns a UserserviceService with all of the expected middleware wired in.
func New(middleware []Middleware) UserserviceService {
	var svc UserserviceService = NewBasicUserserviceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
