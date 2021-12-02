package user

import (
	"fmt"

	"go-api-example/model"
	"go-api-example/pkg/errno"
	"go-api-example/util"

	. "go-api-example/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// 创建用户
func CreateDemo(c *gin.Context) {
	log.Info("创建用户", lager.Data{"X-Request-Id": util.GetReqID(c)})

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

// @Summary 插入一个用户到数据库
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func CreateDataBase(c *gin.Context) {
	log.Info("插入用户", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 参数校验
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 加密密码
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 插入用户
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)
}
