package user

import (
	"fmt"

	"go-api-example/pkg/errno"

	. "go-api-example/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 增加用户
func Create(c *gin.Context) {
	var r CreateRequest

	var err error

	// Bind: 检查Content-Type类型，将消息体作为指定的格式解析到 Go struct 变量中
	// apiserver采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 取param参数
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	// 取query参数
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	// 取请求头Content-Type
	contentType := c.GetHeader("Content-Type")

	log.Infof("Header Content-Type: %s", contentType)
	log.Debugf("username is %s, password is %s", r.Username, r.Password)

	if r.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("用户名为空: xx.xx.xx.xx")).Add("这是自己添加的信息"), nil)
		log.Errorf(err, "Get an error")
		return
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)
}
