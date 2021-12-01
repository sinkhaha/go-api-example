package user

import (
	. "go-api-example/handler"
	"go-api-example/model"
	"go-api-example/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 根据用户名获取用户详情
func Get(c *gin.Context) {
	username := c.Param("username")

	user, err := model.GetUser(username)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
