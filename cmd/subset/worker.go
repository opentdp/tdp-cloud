package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/service"
)

var workerAct string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "子节点管理",
	Run: func(cmd *cobra.Command, args []string) {
		service.Control("worker", workerAct)
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
