package model

import "time"

type SignInfo struct {
	Date     time.Time `gorm:"type:datetime;column:date"`
	User_Seq string    `gorm:"type:varchar(50);column:user_seq"`
	User     User      `gorm:"ForeignKey:User_Seq;"`
}

func (SignInfo) TableName() string {
	return "t_sign"
}
