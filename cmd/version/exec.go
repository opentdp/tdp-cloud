package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
)

func Execute(cmd *cobra.Command, params []string) {

	fmt.Println("Version", args.Version)
	fmt.Println("Build", args.BuildVersion)

}
