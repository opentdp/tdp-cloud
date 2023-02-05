package version

import (
	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use:   "version",
	Short: "显示版本号",
	Run:   Execute,
}

func WithCli() *cobra.Command {

	cli.Flags().BoolP("help", "p", false, "查看帮助")
	cli.Flags().MarkHidden("help")

	return cli

}
