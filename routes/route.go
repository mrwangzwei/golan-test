package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self-test/app/controller"
	"self-test/app/middleware"
	"self-test/config"
)

func InitRoutes() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	fmt.Println("web server running...")
	//中间件
	router.Use(middleware.Recover())
	//路由开始
	userInfo := router.Group("user")
	{
		userInfo.POST("/test", controller.UserInfo.FindUserInfo)
	}
	return router.Run(config.Conf.WebListen)
}
