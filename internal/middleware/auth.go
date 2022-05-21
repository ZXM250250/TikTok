package middleware

import (
	"TikTok/internal/log"
	"TikTok/pkg/common/response"
	"TikTok/pkg/common/secure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Query("token")

		if token == "" {
			c.JSON(http.StatusOK,
				response.CommResponse{
					StatusCode: response.FailureCode,
					StatusMsg:  response.FailureMsgToken,
				})
			c.Abort()
			return
		}
		log.Info("get token=", token)
		account, err := secure.ParserToken(token)

		if err != nil {
			c.JSON(http.StatusOK, response.CommResponse{
				StatusCode: response.FailureCode,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("account", account)

	}
}
