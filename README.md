# go-api-example

## 技术栈
gin + gorm

## 功能
* API服务器状态检查
* 登录用户
* 新增用户
* 删除用户
* 更新用户
* 获取指定用户的详细信息
* 获取用户列表

# 启动http服务
```bash
# 初始化go mod
go mod init go-api-example

# 下载gin
go get -u github.com/gin-gonic/gin

# 下载gopsutil包
go get -u github.com/shirou/gopsutil

# 启动服务
go run main.go
或
go build main.go
 ./main 

# 访问
curl http://localhost:8080/home/health

```

## 测试创建用户
```bash
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user
# 响应  没有传入参数返回 errno.ErrBind错误
{"error":{"Code":10002,"Message":"Error occurred while binding the request body to the struct."}}


curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin"}'
# 响应
{"code":10001,"message":"password is empty"}


curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"password":"admin"}'
# 响应
{"code":20102,"message":"用户不存在 This is add message."}

curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin","password":"admin"}'
# 响应
{"code":0,"message":"OK"}
```