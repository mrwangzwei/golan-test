package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"self-test/config"
	"self-test/dao/mysql"
	"self-test/routes"
)

var Commands = []*cobra.Command{
	Test,
	WebServer,
}

var (
	WebServer = &cobra.Command{
		Use:   "start_web_server",
		Short: "webServer start",
		Run:   startWebServer,
	}

	Test = &cobra.Command{
		Use:   "test",
		Short: "test command",
		Run:   start,
	}
)

func startWebServer(c *cobra.Command, args []string) {
	//初始化mysql
	mysql.InitMysql(cfg)
	err := routes.InitRoutes()
	if err != nil {
		panic(err)
	}
}

func start(c *cobra.Command, args []string) {
	fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println(config.Conf.ServerName)
}
