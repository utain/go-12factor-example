package cmd

import "github.com/spf13/cobra"

import "fmt"

var (
	versionCMD = &cobra.Command{
		Use:   "version",
		Short: "check version of command",
		Long:  `check version of command`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Go Example V1.0.0")
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCMD)
}
