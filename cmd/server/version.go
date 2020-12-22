package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// GitCommit inject by -ldflags
// GIT_COMMIT=git rev-list -1 HEAD && go build -ldflags "-X main.GitCommit=$GIT_COMMIT"
var GitCommit string

// Version of module from go.mod
var Version string

//BuildDate time of build
var BuildDate string

var (
	versionCMD = &cobra.Command{
		Use:   "version",
		Short: "check version of command",
		Long:  `check version of command`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("App Version:", Version)
			fmt.Println("Code Version:", GitCommit)
			fmt.Println("Build Date:", GitCommit)
			fmt.Println("Go Version:", runtime.Version())
			fmt.Println("GOOS:", runtime.GOOS)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCMD)
}
