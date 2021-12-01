package handler

import (
	"net/http"

	"go-api-example/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 响应的固定格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 响应函数
func SendResponse(c *gin.Context, err error, data interface{}) {
	// 解析err得到code和message
	code, message := errno.DecodeErr(err)

	// 都是200状态
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
