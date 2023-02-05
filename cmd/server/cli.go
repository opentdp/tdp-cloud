package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var svc string

var cli = &cobra.Command{
	Use:   "server",
	Short: "启动服务端",
	Run:   Execute,
}

func WithCli() *cobra.Command {

	cli.Flags().BoolP("help", "p", false, "查看帮助")
	cli.Flags().MarkHidden("help")

	cli.Flags().StringVarP(&svc, "service", "s", "", "管理系统服务")

	cli.Flags().StringP("listen", "l", ":7800", "服务端监听的IP地址和端口")
	cli.Flags().StringP("dsn", "d", "server.db", "数据源名称，支持MySQL和SQLite")

	viper.BindPFlag("server.listen", cli.Flags().Lookup("listen"))
	viper.BindPFlag("server.dsn", cli.Flags().Lookup("dsn"))

	return cli

}
