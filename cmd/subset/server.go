package subset

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/helper/strutil"
	"tdp-cloud/service/server"
)

var serverSvc string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务端",
	Run:   serverRun,
}

func serverRun(cmd *cobra.Command, params []string) {

	var err error
	var svc = server.Service()

	log.Println(strutil.FirstUpper(serverSvc), "service tdp-server ...")

	switch serverSvc {
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

func WithServer() *cobra.Command {

	serverCmd.Flags().BoolP("help", "p", false, "查看帮助")
	serverCmd.Flags().MarkHidden("help")

	serverCmd.Flags().StringVarP(&serverSvc, "service", "s", "", "管理系统服务")

	serverCmd.Flags().StringP("listen", "l", ":7800", "服务端监听的IP地址和端口")
	serverCmd.Flags().StringP("dsn", "d", "server.db", "数据源名称，支持MySQL和SQLite")

	viper.BindPFlag("server.listen", serverCmd.Flags().Lookup("listen"))
	viper.BindPFlag("server.dsn", serverCmd.Flags().Lookup("dsn"))

	return serverCmd

}
