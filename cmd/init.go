package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
)

var rcmd = &cobra.Command{
	Use:     "tdp-cloud",
	Short:   "TDP Cloud",
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	cobra.OnInitialize(initViper)

	rcmd.PersistentFlags().StringVarP(&args.ConfigFile, "config", "c", "", "配置文件路径")

}

func initViper() {

	viper.AutomaticEnv()

	if args.ConfigFile == "" {
		if os.Getenv("TDP_DEBUG") != "" {
			log.Println("Configuration file ignored")
		}
		return
	}

	viper.SetConfigFile(args.ConfigFile)
	viper.SafeWriteConfigAs(args.ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}

}
