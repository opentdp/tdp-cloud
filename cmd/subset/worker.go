package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/service"
)

var workerAct string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "子节点管理",
	Run: func(cmd *cobra.Command, rq []string) {
		args.SubCommand.Name = cmd.Name()
		args.SubCommand.Action = serverAct
		service.Control(args.SubCommand.Name, serverAct)
	},
}

func WithWorker() *cobra.Command {

	workerCmd.Flags().BoolP("help", "h", false, "查看帮助")
	workerCmd.Flags().MarkHidden("help")

	workerCmd.Flags().StringVarP(&workerAct, "service", "s", "", "管理系统服务")
	workerCmd.Flags().StringP("remote", "r", "", "注册地址 (e.g. ws://{domain}/wsi/{appid}/worker)")

	viper.BindPFlag("worker.remote", workerCmd.Flags().Lookup("remote"))

	return workerCmd

}
