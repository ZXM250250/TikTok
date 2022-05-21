package service

import (
	"TikTok/internal/dao"
	"TikTok/internal/log"
	"TikTok/internal/model"
	"TikTok/pkg/common/response"
	"TikTok/pkg/common/secure"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = dao.GetDB()
}
func Register(username string, password string, res *response.Response) (err error) {
	salt, err := secure.HashAndSalt(password)

	account := user.Account{Username: username, Password: string(salt)}

	result := db.Create(&account)

	res.UserId = account.ID
	res.Token, err = secure.GenerateToken(account)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	if result.Error != nil {
		res.StatusMsg = "用户名已经存在了"
		return result.Error
	}

	res.StatusCode = 0
	res.StatusMsg = "注册成功"
	return nil
}

func Login(username string, password string, res *response.Response) (err error) {

	var account user.Account
	result := db.Where("username=?", username).First(&account)
	res.UserId = account.ID

	if result.Error == gorm.ErrRecordNotFound {
		res.StatusMsg = "未找到用户相关的信息"
		return
	}
	res.Token, err = secure.GenerateToken(account)
	if !secure.ComparePasswords(account.Password, password) {
		res.StatusMsg = "登陆失败,密码不正确"
		return
	}
	res.StatusMsg = "登陆成功"
	res.StatusCode = 0

	return
}

func GetUserInfo(id string, info *response.UserInfo) (err error) {
	var account user.Account
	result := db.Where("id=?", id).First(&account)
	if result.Error != nil {
		log.Errorf(result.Error)
		return result.Error

	}
	info.User.Username = account.Username
	info.User.ID = account.ID
	info.User.FollowCount = account.FollowCount
	info.User.FollowerCount = account.FollowerCount
	info.User.IsFollow = account.IsFollow

	return nil

}
