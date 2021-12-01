package user

import (
	. "go-api-example/handler"
	"go-api-example/model"
	"go-api-example/pkg/auth"
	"go-api-example/pkg/errno"
	"go-api-example/pkg/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// 校验密码
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// 生成token
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	// 响应token
	SendResponse(c, nil, model.Token{Token: t})
}
