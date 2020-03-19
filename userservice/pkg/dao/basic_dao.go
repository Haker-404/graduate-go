package dao

import (
	"awesomeProject/userservice/pkg/model"
	"github.com/jinzhu/gorm"
	"time"
)

type BasicDao interface {
	GetUserInfo(seq string, pwd ...string) (model.User, model.Resp)
	InsertSign(seq string) error
}
type UserDao struct {
}

var DB = &gorm.DB{}

func (UserDao) GetUserInfo(seq string, pwd ...string) (user model.User, message model.Resp) {
	user = model.User{}
	if pwd != nil {
		DB.Where("seq = ? AND password = ?", seq, pwd).Find(&user)
	} else {
		DB.First(&user, seq)
	}
	if user.Seq == "" {
		message.Msg = "查询失败，无此用户"
		message.Succ = "false"
	} else {
		message.Msg = "查询成功"
		message.Succ = "true"
	}
	return
}
func (UserDao) InsertSign(seq string) error {
	signInfo := &model.SignInfo{
		Date:     time.Now(),
		User_Seq: seq,
	}
	db := DB.Create(&signInfo)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func OpenDB() {
	db, err := gorm.Open("mysql", "root:111111@tcp(127.0.0.1:3306)/graduatedDesign?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	DB = db
	DB.SingularTable(true)
	db.LogMode(true)
	//defer DB.Close()
}
