package router

import (
	"go-api-example/handler/sd"
	"go-api-example/handler/user"
	"go-api-example/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 注册全局中间件
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404处理
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "路径错误")
	})

	// 注册路由
	// home分组下有 /health  /disk  /cpu  /ram等4个路由
	homeRouteGroup := g.Group("/home")
	{
		homeRouteGroup.GET("/health", sd.HealthCheck)
		homeRouteGroup.GET("/disk", sd.DiskCheck)
		homeRouteGroup.GET("/cpu", sd.CPUCheck)
		homeRouteGroup.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("/v1/user")
	{
		// 测试创建用户
		u.POST("/:username", user.CreateDemo)

		// 创建用户
		u.POST("", user.CreateDataBase)
		// 删除用户
		u.DELETE("/:id", user.Delete)
		// 更新用户
		u.PUT("/:id", user.Update)
		// 用户列表(分页)
		u.GET("", user.List)
		// 获取用户详情
		u.GET("/:username", user.Get)
	}

	return g
}
