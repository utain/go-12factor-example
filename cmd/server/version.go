package main

import "github.com/spf13/cobra"

import "fmt"

// GitCommit inject by -ldflags
// GIT_COMMIT=git rev-list -1 HEAD && go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
var GitCommit string

// Version of module from go.mod
var Version string

var (
	versionCMD = &cobra.Command{
		Use:   "version",
		Short: "check version of command",
		Long:  `check version of command`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Server Version:", Version)
			fmt.Println("Git Commit:", GitCommit)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCMD)
}
