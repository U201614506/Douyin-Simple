package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
 

var db *gorm.DB

func Connection() (*gorm.DB) {
	dsn := "huang:root@tcp(49.235.116.64:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=  nil {
		panic(err)
	}
	return db
}
