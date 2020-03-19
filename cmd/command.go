package cmd

import (
	"github.com/spf13/cobra"
	"self-test/config"
	"self-test/crontask"
	"self-test/dao/mysql"
	"self-test/dao/redis"
	"self-test/routes"
)

func InitServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start_web_server",
		Short: "webServer start",
		Run:   startWebServer,
	}
	config.Conf = config.NewDefaultConfig()
	cmd.Flags().StringVar(&config.Conf.ConfigPath, "config", "./config.yml", "path to the config file")
	return cmd
}

func startWebServer(c *cobra.Command, args []string) {
	err := config.Conf.LoadConfigFile()
	if err != nil {
		panic(err)
	}
	//初始化mysql
	mysql.InitMysql()

	//初始化redis
	redis.Init()

	//开启定时任务
	go crontask.Run()

	err = routes.InitRoutes()
	if err != nil {
		panic(err)
	}
}
