package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/service"
)

var serverAct string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "服务端管理",
	Run: func(cmd *cobra.Command, args []string) {
		service.Control("server", serverAct)
	},
}

func WithServer() *cobra.Command {

	serverCmd.Flags().BoolP("help", "h", false, "查看帮助")
	serverCmd.Flags().MarkHidden("help")

	serverCmd.Flags().StringVarP(&serverAct, "service", "s", "", "管理系统服务")

	serverCmd.Flags().StringP("listen", "l", ":7800", "服务端监听的IP地址和端口")
	serverCmd.Flags().StringP("dsn", "d", "server.db", "数据源名称，支持MySQL和SQLite")

	viper.BindPFlag("server.listen", serverCmd.Flags().Lookup("listen"))
	viper.BindPFlag("server.dsn", serverCmd.Flags().Lookup("dsn"))

	return serverCmd

}
