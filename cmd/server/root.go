package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-example",
		Short: "go-example is example golang web service using gin and gorm",
		Long:  `go-example is example golang web service using gin and gorm`,
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
