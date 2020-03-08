package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"self-test/config"
)

var (
	cfg *config.ServerConf
	rootCmd = &cobra.Command{
		Use:   "wzw",
		Short: "this is a cadre",
	}
)

func init() {
	//加载配置
	cfg = config.InitConfigPath(rootCmd)
	err := cfg.LoadConfigFile()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
		return
	}
}

func Start() {
	rootCmd.AddCommand(Commands...)
	if len(os.Args) > 1 {
		if err := rootCmd.Execute(); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
}
