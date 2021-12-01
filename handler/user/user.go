package user

// 存放user业务所有接口 接收json消息体 的struct

// 创建用户接口的请求体
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 创建接口的响应体
type CreateResponse struct {
	Username string `json:"username"`
}
