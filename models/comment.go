package models

import "time"

type Comment struct {
	Id         int64     `gorm:"type:int(11);column:id;primary_key"`
	VideoId     string    `gorm:"type:varchar(255);column:video_id"`
	UserId     string    `gorm:"type:varchar(255);column:user_id"`
	Content    string    `gorm:"type:varchar(255);column:content"`
	Delete         int64     `gorm:"type:int(11);column:delete"`
	CreatedAt time.Time `gorm:"type:timestamp;column:create_time"`
}	

func (Comment) TableName() string {
	return "comment"
}