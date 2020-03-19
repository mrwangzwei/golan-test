package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wzw",
		Short: "this is a cadre",
	}
)

func Start() {
	rootCmd.AddCommand(InitServerCommand())
	if len(os.Args) > 1 {
		if err := rootCmd.Execute(); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
}
