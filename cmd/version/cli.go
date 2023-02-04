package version

import (
	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use:   "version",
	Short: "显示版本号",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func WithCli() *cobra.Command {

	return cli

}
