package middleware

import (
	"go-api-example/handler"
	"go-api-example/pkg/errno"
	"go-api-example/pkg/token"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件 jwt
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)

			c.Abort()

			return
		}

		c.Next()
	}
}
