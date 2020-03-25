package dao

import (
	"awesomeProject/userservice/pkg/model"
	"github.com/jinzhu/gorm"
	"time"
)

type BasicDao interface {
	GetUserInfo(seq string, pwd ...string) (model.User, model.Resp)
	InsertSign(seq string) error
	GetUserList(date time.Time, isSigned bool) (signInfoList []model.SignInfo, message model.Resp)
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
func (UserDao) GetUserList(date time.Time, isSigned bool) (signInfoList []model.SignInfo, message model.Resp) {
	switch isSigned {
	case true:
		if date.Hour() == 0 && date.Minute() == 0 && date.Second() == 0 {
			t_time := date.Format("2006-01-02")
			DB.Table("t_sign").Where("Date(date) = ?", t_time).Preload("User").Find(&signInfoList)
		} else {
			DB.Table("t_sign").Where("date = ?", date).Preload("User").Find(&signInfoList)
		}
	case false:
		if date.Hour() == 0 && date.Minute() == 0 && date.Second() == 0 {
			t_time := date.Format("2006-01-02")
			DB.Debug().Exec("SELECT * FROM `t_user` WHERE `t_user`.`seq` NOT in ( SELECT `t_sign`.`user_seq` FROM `t_sign` WHERE Date( `date` ) = ? )", t_time)
			// DB.Debug().Exec("SELECT t_user.* FROM t_user where not exists (select t_sign.* from t_sign where t_sign.user_seq=t_user.seq and Date(t_sign.date)!=?)",t_time).Find(&signInfoList)
		} else {
			DB.Debug().Exec("SELECT b.* FROM t_user b where not exists (select a.* from t_sign a where a.user_seq=b.seq and Date(a.date)!=?)", date).Find(&signInfoList)
		}
	}
	message.Msg = "查询成功"
	message.Succ = "true"
	return
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
