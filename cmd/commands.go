package cmd

import (
	"github.com/spf13/cobra"
	"self-test/cmd/commands"
	"self-test/route"
)

var Commands = []*cobra.Command{
	commands.Test,
	route.WebServer,
}

