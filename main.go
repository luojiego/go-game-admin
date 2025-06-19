package main

import (
	"log"
	"net/http"

	"management/config"
	"management/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	router := gin.Default()

	router.Static("/dist", "./dist")
	router.Static("/plugins", "./plugins")
	router.LoadHTMLGlob("templates/**/*")

	// 创建路由管理器
	routeManager := handlers.NewRouteManager(router)

	// 注册API路由
	handlers.RegisterAPIRoutes(router)

	// 加载动态路由
	if err := routeManager.LoadRoutes(); err != nil {
		log.Fatal("Failed to load routes:", err)
	}

	// 添加路由管理界面
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"Title": "Route Management",
		})
	})

	router.Run(":9080")
}
