package model

import (
	"TikTok/internal/dao"
	"TikTok/internal/log"
	"github.com/golang-jwt/jwt"
)

type Account struct {
	ID                 uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username           string `gorm:"unique" json:"name"`
	Password           string `json:"-"`
	FollowCount        uint   `json:"follow_count"`
	FollowerCount      uint   `json:"follower_count"`
	IsFollow           bool   `json:"is_follow"`
	jwt.StandardClaims `gorm:"-" json:"-"`
}

func init() {
	err := dao.GetDB().AutoMigrate(&Account{}, &Video{})
	if err != nil {
		log.Errorf("建表失败错误", err)
		return
	}
}
