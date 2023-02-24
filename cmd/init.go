package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/initd"
)

var rcmd = &cobra.Command{
	Use:     "tdp-cloud",
	Short:   "TDP Cloud",
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	// 延迟执行

	cobra.OnInitialize(
		initd.Viper, initd.Dataset, initd.Logman,
	)

	// 全局参数

	rcmd.PersistentFlags().StringVarP(&args.ConfigFile, "config", "c", "", "配置文件路径")
	rcmd.PersistentFlags().StringP("datadir", "", "var/data", "数据存储目录")
	rcmd.PersistentFlags().StringP("logdir", "", "var/log", "日志存储目录")

	viper.BindPFlag("dataset.dir", rcmd.PersistentFlags().Lookup("datadir"))
	viper.BindPFlag("logger.dir", rcmd.PersistentFlags().Lookup("logdir"))

}
