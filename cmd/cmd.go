package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gorm",
		Short: "gorm is a very fast static site generator",
		Long:  `gorm is command line`,
	}
)

// Execute root command
func Execute() {
	rootCmd.Execute()
}
