package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"go-api-example/config"
	"go-api-example/model"
	v "go-api-example/pkg/version"
	"go-api-example/router"
	"go-api-example/router/middleware"

	"github.com/gin-gonic/gin"

	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	// cfg变量值从命令行flag传入，可以传值，比如 ./main -c config.yaml，也可以为空，如果为空会默认读取 conf/config.yaml
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
	// 接收命令行中-v参数所带的版本信息
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()

	// 运行程序时带上-v即可查看版本信息
	if *version {
		v := v.Get()

		// 格式化版本信息
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 建立数据库连接
	model.DB.Init()
	// 关闭数据库连接
	defer model.DB.Close()

	// 读取配置runmode的值，设置gin模式，gin 有 3 种运行模式：debug、release 和 test，debug模式会打印很多debug信息
	gin.SetMode(viper.GetString("runmode"))

	// 开个协程循环打印环境，可验证修改runmode的值后会监听配置进行热更新
	go func() {
		for {
			log.Infof("current mode: %s", viper.GetString("runmode"))
			time.Sleep(5 * time.Second)
		}
	}()

	// create a gin engine
	g := gin.New()

	// middlewares := []gin.HandlerFunc{}

	router.Load(
		g,

		// 全局中间件 最终会调用g.Use()加载该中间件
		middleware.Logging(),
		middleware.RequestId(),
	)

	// 开启个协程去请求/sd/health路由
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("router没反应，超时", err)
		}
		log.Info("router成功加载")
	}()

	addr := viper.GetString("addr")
	log.Infof("listening %s", addr)
	log.Info(http.ListenAndServe(addr, g).Error())
}

func pingServer() error {
	// 获取配置max_ping_count变量的值
	retryCount := viper.GetInt("max_ping_count")

	// 重试2次，状态码200则可用
	for i := 0; i < retryCount; i++ {
		res, err := http.Get(viper.GetString("url") + "/home/health")
		if err == nil && res.StatusCode == 200 {
			return nil
		}

		log.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("健康检查失败")
}
