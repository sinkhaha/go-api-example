package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// 构造id
func GenShortId() (string, error) {
	return shortid.Generate()
}

// 从请求头X-Request-Id获取请求id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
