package service

import "errors"

type Service interface {
	GetName(userid string) string
}
type UserService struct {
}
func (userservice UserService)GetName(userid string) string{
	if userid==""{
		return errors.New("id为空").Error()
	}
	return userid
}