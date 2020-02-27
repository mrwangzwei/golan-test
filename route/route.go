package route

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"self-test/app/controller"
	"self-test/app/middleware"
	"self-test/config"
)

var (
	WebServer = &cobra.Command{
		Use:   "start_web_server",
		Short: "webServer start",
		Run:   startWebServer,
	}
)

func startWebServer(cmd *cobra.Command, args []string) {
	err := InitRoutes()
	if (err != nil) {
		panic(err)
	}
}

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
