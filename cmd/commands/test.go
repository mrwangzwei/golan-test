package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"self-test/config"
)

var (
	Test = &cobra.Command{
		Use:   "test",
		Short: "test command",
		Run:   start,
	}
)

func start(cmd *cobra.Command, args []string){
	fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println(config.Conf.ServerName)
}
