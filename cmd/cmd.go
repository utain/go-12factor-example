package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-example",
		Short: "go-example is a very fast static site generator",
		Long:  `go-example is command line`,
	}
)

// Execute root command
func Execute() {
	rootCmd.Execute()
}
