package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/initd"
	"tdp-cloud/cmd/subset"
)

var rcmd = &cobra.Command{
	Use:     "tdp-cloud",
	Short:   args.AppName,
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	// 延迟执行

	cobra.OnInitialize(
		initd.Viper, initd.Dataset, initd.Logger,
	)

	// 全局参数

	rcmd.PersistentFlags().StringVarP(&initd.ViperFile, "config", "c", "config.yml", "config file path")

}

func Execute() {

	rcmd.AddCommand(
		subset.WithServer(), subset.WithWorker(), subset.WithUpdate(),
	)

	if err := rcmd.Execute(); err != nil {
		os.Exit(1)
	}

}
