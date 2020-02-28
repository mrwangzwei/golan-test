package cmd

import (
	"github.com/spf13/cobra"
	"self-test/cmd/commands"
)

var Commands = []*cobra.Command{
	commands.Test,
	commands.WebServer,
}

