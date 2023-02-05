package subset

import (
	"fmt"

	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本号",
	Run: func(cmd *cobra.Command, params []string) {
		fmt.Println("Version", args.Version)
		fmt.Println("Build", args.BuildVersion)
	},
}

func WithVersion() *cobra.Command {

	versionCmd.Flags().BoolP("help", "p", false, "查看帮助")
	versionCmd.Flags().MarkHidden("help")

	return versionCmd

}
