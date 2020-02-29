package routes

import (
	"github.com/gin-gonic/gin"
	"self-test/app/controller"
	"self-test/app/middleware"
	"self-test/config"
)

func InitRoutes() error {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//中间件
	router.Use(middleware.Recover())
	//路由开始
	userInfo := router.Group("user")
	{
		userInfo.POST("/test", controller.FindUserInfo)
	}
	return router.Run(config.Conf.WebListen)
}