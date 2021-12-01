package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 增加请求id的中间件
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		// 不存在则创建一个
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// 设置到上下文
		c.Set("X-Request-Id", requestId)

		// 设置请求头
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
