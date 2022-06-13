package models

type Favorite struct {
	Id      int64  `gorm:"type:int(11);column:id;primary_key"`
	ViedoId string `gorm:"type:varchar(255);column:video_id;"`
	UserId  string `gorm:"type:varchar(255);column:user_id;"`
	Delete  int64  `gorm:"type:int(11);column:delete;omitempty;default:0"`
}

func (Favorite) TableName() string {
	return "favorite"
}