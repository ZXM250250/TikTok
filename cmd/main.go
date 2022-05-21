package main

import (
	"TikTok/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.InitRouter(r)
	r.Run()

}
