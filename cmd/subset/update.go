package subset

import (
	"github.com/open-tdp/go-helper/upgrade"
	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "下载更新包",
	Long:  "下载完成后，请重启服务完成更新",
	Run: func(cmd *cobra.Command, rq []string) {
		ExecUpdate()
	},
}

func WithUpdate() *cobra.Command {

	updateCmd.Flags().BoolP("help", "h", false, "查看帮助")
	updateCmd.Flags().MarkHidden("help")

	return updateCmd

}

func ExecUpdate() error {

	err := upgrade.Apply(&upgrade.RequesParam{
		UpdateUrl: args.UpdateUrl,
		Version:   args.Version,
	})

	return err

}
