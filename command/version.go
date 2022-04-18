package command

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/utain/go/example/internal/version"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print application version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Version:", version.Version)
			fmt.Println("Built:", version.Time)
			fmt.Println("Commit:", version.Commit)
			fmt.Println("Runtime:", fmt.Sprintf("%s %s %s", runtime.Version(), runtime.GOARCH, runtime.Compiler))
			return nil
		},
	})
}
