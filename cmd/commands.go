package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"self-test/config"
	"self-test/crontask"
	"self-test/dao/mysql"
	"self-test/dao/redis"
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
	config.Conf = config.NewDefaultConfig()
	c.Flags().StringVar(&config.Conf.ConfigPath, "config", "./config.yml", "path to the config file")
	err := config.Conf.LoadConfigFile()
	if err != nil {
		panic(err)
	}
	//初始化mysql
	mysql.InitMysql()

	//初始化redis
	redis.Init()

	//开启定时任务
	crontask.Run()

	err = routes.InitRoutes()
	if err != nil {
		panic(err)
	}
}

func start(c *cobra.Command, args []string) {
	fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println(config.Conf.ServerName)
}
