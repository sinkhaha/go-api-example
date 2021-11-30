# go-api-example
## 账号系统功能
* API服务器状态检查
* 登录用户
* 新增用户
* 删除用户
* 更新用户
* 获取指定用户的详细信息
* 获取用户列表

技术栈：gin + gorm
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
