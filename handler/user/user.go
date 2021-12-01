package user

// 存放user业务所有接口 接收json消息体 的struct

import (
	"go-api-example/model"
)

// 创建用户接口的请求体
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 创建接口的响应体
type CreateResponse struct {
	Username string `json:"username"`
}

// 列表请求体
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// 列表响应体
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
