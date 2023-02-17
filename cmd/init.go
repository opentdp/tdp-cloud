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

	cobra.OnInitialize(initHook) // 初始化后触发钩子

	rcmd.PersistentFlags().StringVarP(&args.ConfigFile, "config", "c", "", "配置文件路径")

}

func initHook() {

	initViper()
	initLogger()

}

func initViper() {

	viper.AutomaticEnv()

	if args.ConfigFile == "" {
		if os.Getenv("TDP_DEBUG") != "" {
			log.Println("Configuration file ignored.")
		}
		return
	}

	viper.SetConfigFile(args.ConfigFile)
	viper.SafeWriteConfigAs(args.ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}

}

func initLogger() {

	log.SetPrefix("[TDP] ")

	if logPath := viper.GetString("logger.directory"); logPath != "" {
		logFile, err := os.OpenFile(logPath+"/output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Failed to", err)
			log.Println("Fallback to using standard output.")
		} else {
			log.SetOutput(logFile)
		}
	}

}
