package models

import "time"

type Video struct {
	Id         int64     `gorm:"type:int(11);column:id;primary_key"`
	Author     string    `gorm:"type:varchar(255);column:author"`
	PlayUrl    string    `gorm:"type:varchar(255);column:play_url"`
	CoverUrl   string    `gorm:"type:varchar(255);column:cover_url"`
	Title      string    `gorm:"type:varchar(255);column:title"`
	CreatedAt time.Time `gorm:"type:timestamp;column:create_time"`
}	

func (Video) TableName() string {
	return "video"
}