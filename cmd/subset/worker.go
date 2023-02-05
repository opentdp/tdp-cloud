package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/service/worker"
)

var workerSvc string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "注册子节点",
	Run: func(cmd *cobra.Command, params []string) {
		switch workerSvc {
		case "install":
			worker.Service().Install()
		case "uninstall":
			worker.Service().Uninstall()
		case "":
			worker.Service().Run()
		}
	},
}

func WithWorker() *cobra.Command {

	workerCmd.Flags().BoolP("help", "p", false, "查看帮助")
	workerCmd.Flags().MarkHidden("help")

	workerCmd.Flags().StringVarP(&workerSvc, "service", "s", "", "管理系统服务")

	workerCmd.Flags().StringP("remote", "r", "", "注册地址 (e.g. ws://{domain}/wsi/{appid}/worker)")

	viper.BindPFlag("worker.remote", workerCmd.Flags().Lookup("remote"))

	return workerCmd

}
