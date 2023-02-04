package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cli = &cobra.Command{
	Use:   "server",
	Short: "启动服务端",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func WithCli() *cobra.Command {

	cli.Flags().StringP("listen", "l", ":7800", "服务端监听的IP地址和端口")
	cli.Flags().StringP("dsn", "d", "server.db", "数据源名称，支持MySQL和SQLite")

	viper.BindPFlag("server.listen", cli.Flags().Lookup("listen"))
	viper.BindPFlag("server.dsn", cli.Flags().Lookup("dsn"))

	return cli

}
