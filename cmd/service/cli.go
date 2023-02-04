package service

import (
	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use:   "service",
	Short: "安装或卸载服务",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func WithCli() *cobra.Command {

	cli.Flags().StringP("install", "i", "", "安装服务")
	cli.Flags().StringP("uninstall", "u", "", "卸载服务")

	return cli

}
