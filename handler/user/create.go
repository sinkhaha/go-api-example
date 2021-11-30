package user

import (
	"fmt"
	"net/http"

	"go-api-example/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 增加用户
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error

	// 判断参数是否正常传递
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("用户名为空: xx.xx.xx.xx")).Add("这是自己添加的信息")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
