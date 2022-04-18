package command

import (
	"github.com/spf13/cobra"
)

type ServerOpts struct {
	Port uint16
}

type Versions struct {
	Version string
	Runtime string
}

var root = cobra.Command{
	Use: "go-example",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var port uint16

func Execute() error {
	return root.Execute()
}
