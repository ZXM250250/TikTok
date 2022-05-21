package controllers

import (
	user "TikTok/internal/model"
	"github.com/gin-gonic/gin"
)

func publish(c *gin.Context) {
	account, _ := c.Get("account")
	user := account.(user.Account)

}

func publishList(c *gin.Context) {

}
