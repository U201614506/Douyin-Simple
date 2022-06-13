package models

type User struct {
	Id       int64  `gorm:"type:int(11);column:id;primary_key"`
	UserName string `gorm:"type:varchar(255);column:username"`
	Password string `gorm:"type:varchar(255);column:password"`
}

func (User) TableName() string {
	return "user"
}