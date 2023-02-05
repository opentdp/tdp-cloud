package subset

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/service/worker"
)

var workerSvc string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "注册子节点",
	Run:   workerRun,
}

func workerRun(cmd *cobra.Command, params []string) {

	var err error
	var svc = worker.Service()

	log.Println(strutil.FirstUpper(workerSvc), "service tdp-worker ...")

	switch workerSvc {
	case "install":
		err = svc.Install()
	case "uninstall":
		err = svc.Uninstall()
	case "start":
		err = svc.Start()
	case "status":
		_, err = svc.Status()
	case "stop":
		err = svc.Stop()
	case "restart":
		err = svc.Restart()
	case "":
		err = svc.Run()
	}

	if err != nil {
		log.Println(err)
	}

}

func WithWorker() *cobra.Command {

	workerCmd.Flags().BoolP("help", "p", false, "查看帮助")
	workerCmd.Flags().MarkHidden("help")

	workerCmd.Flags().StringVarP(&workerSvc, "service", "s", "", "管理系统服务")

	workerCmd.Flags().StringP("remote", "r", "", "注册地址 (e.g. ws://{domain}/wsi/{appid}/worker)")

	viper.BindPFlag("worker.remote", workerCmd.Flags().Lookup("remote"))

	return workerCmd

}
