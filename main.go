package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go-api-example/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin engine
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g, middlewares...)

	// 开启个协程去请求/sd/health路由
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("router没反应，超时", err)
		}
		log.Print("router成功加载")
	}()

	log.Printf("listening %s", "8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

func pingServer() error {
	// 重试2次，状态码200则可用
	for i := 0; i < 2; i++ {
		res, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && res.StatusCode == 200 {
			return nil
		}

		log.Print("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("不能请求通")
}
