package commands

import (
	"github.com/spf13/cobra"
	"self-test/routes"
)

var (
	WebServer = &cobra.Command{
		Use:   "start_web_server",
		Short: "webServer start",
		Run:   startWebServer,
	}
)

func startWebServer(c *cobra.Command, args []string) {
	err := routes.InitRoutes()
	if (err != nil) {
		panic(err)
	}
}
