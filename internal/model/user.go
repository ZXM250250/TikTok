package model

import (
	"TikTok/internal/dao"
	"TikTok/internal/log"
	"github.com/golang-jwt/jwt"
)

type Account struct {
	ID                 uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Username           string `gorm:"unique" json:"name"`
	Password           string `json:"-"`
	FollowCount        uint64 `json:"follow_count"`
	FollowerCount      uint64 `json:"follower_count"`
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
