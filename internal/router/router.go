package router

import (
	"TikTok/api/controllers"
	"TikTok/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	groupUser := r.Group("/douyin/user")
	{
		groupUser.POST("/register/", controllers.Register)
		groupUser.POST("/login/", controllers.Login)
	}
	group := r.Group("/douyin", middleware.JWTAuth())
	{
		group.GET("/user/", controllers.GetUserInfo)
		group.POST("/publish/action/", controllers.Publish)

	}

}
