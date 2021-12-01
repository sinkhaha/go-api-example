package user

import (
	"strconv"

	. "go-api-example/handler"
	"go-api-example/model"
	"go-api-example/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 根据id删除用户
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
