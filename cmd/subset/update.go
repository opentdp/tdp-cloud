package subset

import (
	"time"

	"github.com/open-tdp/go-helper/logman"
	"github.com/open-tdp/go-helper/upgrade"
	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
	"tdp-cloud/service"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新程序",
	Run: func(cmd *cobra.Command, rq []string) {
		ExecUpdate()
	},
}

func WithUpdate() *cobra.Command {

	updateCmd.Flags().BoolP("help", "h", false, "查看帮助")
	updateCmd.Flags().MarkHidden("help")

	go AutoUpdate() // 自动更新

	return updateCmd

}

func ExecUpdate() error {

	err := upgrade.Apply(&upgrade.RequesParam{
		UpdateUrl: args.UpdateUrl,
		Version:   args.Version,
	})

	if err == nil {
		logman.Warn("下载完成，请重启服务完成更新")
	}

	return err

}

func AutoUpdate() {

	time.Sleep(15 * time.Second)

	if args.SubCommand.Action == "" {
		logman.Warn("非服务模式，自动更新已关闭")
		return
	}

	for {
		logman.Error("检查更新")
		if err := ExecUpdate(); err != nil {
			logman.Warn("未更新", "info", err)
		} else {
			service.Control(args.SubCommand.Name, "restart")
		}
		time.Sleep(1 * time.Minute)
	}

}
