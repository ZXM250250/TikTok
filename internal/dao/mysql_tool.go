package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	var err error

	_db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/tiktok"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败,error =" + err.Error())
	}

}

func GetDB() *gorm.DB {
	return _db

}
