package model

type User struct {
	Name      string `gorm:"type:varchar(50);column:user_name"`
	Seq       string `gorm:"primary_key;not null;type:varchar(50);unique;column:seq"`
	Photo     []byte `gorm:"type:blob;unique;column:photo"`
	Password  string `gorm:"type:varchar(255);column:password"`
	IsManager bool   `gorm:"type:tinyint(1);column:is_manager"`
}

func (User) TableName() string {
	return "t_user"
}
