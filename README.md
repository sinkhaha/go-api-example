# go-api-example

## 技术栈
gin + gorm



## 文件

```bash
.
├── Makefile             # 编译脚本
├── README.md
├── admin.sh             # 启动脚本 管理启动、重启、停止和查看运行状态等命令
├── conf                 # 配置文件目录
│   └── config.yaml      # 配置文件
├── config               # 处理配置和配置文件的go代码
│   └── config.go
├── db.sql               # mysql脚本
├── docs                 # swagger文档，执行swag init生成
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go-api-example
├── go.mod
├── go.sum
├── handler              # 控制器逻辑
│   ├── handler.go
│   ├── sd               # 健康检查控制器
│   │   └── check.go
│   └── user             # 用户模块控制器
│       ├── create.go    
│       ├── delete.go
│       ├── get.go
│       ├── list.go
│       ├── login.go
│       ├── update.go
│       └── user.go      # 存放用户业务所有接口 接收json消息体 的struct
├── main.go              # go程序唯一入口
├── model                # 数据库相关操作，包括数据库初始化和对表的增删改查
│   ├── init.go          # 初始化和连接数据库
│   ├── model.go         # 放公用的结构体
│   └── user.go          # 用户相关的CURD
├── pkg                  # 引用的包
│   ├── auth             # 认证
│   │   └── auth.go
│   ├── constvar         # 常量
│   │   └── constvar.go
│   ├── errno            # 错误码
│   │   ├── code.go
│   │   └── errno.go
│   ├── token            # jwt
│   │   └── token.go
│   └── version          # 版本包
│       ├── base.go
│       ├── doc.go
│       └── version.go
├── router               # 路由相关处理
│   ├── middleware       # 中间件
│   │   ├── auth.go
│   │   ├── header.go
│   │   ├── logging.go
│   │   └── requestid.go
│   └── router.go        # 路由
├── service              # 业务处理
│   └── user.go 
└── util                 # 工具类函数存放目录
    ├── util.go
    └── util_test.go
```



## 功能
* API服务器状态检查
* 登录用户
* 新增用户
* 删除用户
* 更新用户
* 获取指定用户的详细信息
* 获取用户列表



## 实践

* 配置文件读取
* 日志库的使用
* 数据库的CURD
* 自定义错误信息
* 读取和返回http请求
* 中间件
* api身份验证：jwt
* 给api添加版本功能
* Makefile
* swagger使用
* 单元测试
* api性能分析



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
或
make
./go-api-example

# 访问
curl http://localhost:8080/home/health

```

## 测试创建用户接口
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