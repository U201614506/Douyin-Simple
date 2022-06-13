package models

type Relation struct {
	Id          int64  `gorm:"type:int(11);column:id;primary_key"`
	Follower_Id string `gorm:"type:varchar(255);column:follower_id"`
	Follow_Id   string `gorm:"type:varchar(255);column:follow_id"`
	Delete      string `gorm:"type:varchar(255);column:delete"`
}

func (Relation) TableName() string {
	return "relation"
}