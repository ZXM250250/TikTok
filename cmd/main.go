package main

import (
	"TikTok/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Static("/assets", "assets")
	router.InitRouter(r)
	r.Run()

}
