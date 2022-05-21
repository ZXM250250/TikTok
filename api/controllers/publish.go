package controllers

import (
	"TikTok/internal/log"
	"TikTok/internal/model"
	"TikTok/internal/service"
	"TikTok/pkg/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const BaseUrl = "E:/Projects/Golang/TikTok/assets"
const Video = "/video"
const Cover = "/cover"

func Publish(c *gin.Context) {
	account, _ := c.Get("account")
	user := account.(*model.Account)
	var video model.Video
	file, err := c.FormFile("data")
	res := response.CommResponse{StatusCode: response.FailureCode, StatusMsg: "发生了错误,投稿失败"}
	if err != nil {
		log.Errorf("发生了获取文件的错误", err)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	video.Title = c.Query("title")
	video.FromWho = user.ID
	PlayUrl := BaseUrl + user.Username + Video
	CoverUrl := BaseUrl + user.Username + Cover

	err, video.PlayUrl = service.SaveUploadedFile(file, PlayUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	video.CoverUrl, err = service.GetSnapshot(video.PlayUrl, CoverUrl, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	err = service.Publish(&video)
	if err != nil {
		log.Errorf(err)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.StatusCode = response.SuccessCode
	res.StatusMsg = "投稿成功"
	c.JSON(http.StatusOK, res)

}

func publishList(c *gin.Context) {

}
