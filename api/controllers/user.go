package controllers

import (
	"TikTok/internal/log"
	"TikTok/internal/service"
	"TikTok/pkg/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {

	username := ctx.Query("username")
	password := ctx.Query("password")
	res := response.Response{StatusCode: 1, StatusMsg: "注册失败"}

	if err := service.Register(username, password, &res); err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		log.Errorf(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	res := response.Response{StatusCode: 1, StatusMsg: "登陆失败"}
	if err := service.Login(username, password, &res); err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		log.Errorf(err.Error())

	}
	ctx.JSON(http.StatusOK, res)
}

func GetUserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	var resUserInfo response.UserInfo
	if err := service.GetUserInfo(userId, &resUserInfo); err != nil {
		resUserInfo.StatusMsg = "获取用户信息失败"
		resUserInfo.StatusCode = response.FailureCode
		c.JSON(http.StatusInternalServerError, resUserInfo)
		return
	}

	resUserInfo.StatusMsg = "获取用户信息成功"
	resUserInfo.StatusCode = response.SuccessCode
	c.JSON(http.StatusOK, resUserInfo)

}
