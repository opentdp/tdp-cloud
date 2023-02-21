package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
)

var vipFile string

var rootCmd = &cobra.Command{
	Use:     "tdp-cloud",
	Short:   "TDP Cloud",
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	log.SetPrefix("[TDP] ") // 优先执行

	cobra.OnInitialize(initViper, initDataset, initLogger) // 延迟执行

	// 全局参数

	rootCmd.PersistentFlags().StringVarP(&vipFile, "config", "c", "", "配置文件路径")
	rootCmd.PersistentFlags().StringP("datadir", "", "var/data", "数据存储目录")
	rootCmd.PersistentFlags().StringP("logdir", "", "var/log", "日志存储目录")

	viper.BindPFlag("dataset.dir", rootCmd.PersistentFlags().Lookup("datadir"))
	viper.BindPFlag("logger.dir", rootCmd.PersistentFlags().Lookup("logdir"))

	// 默认参数

	viper.SetDefault("logger.output", false)

}

func initViper() {

	viper.SetEnvPrefix("TDP")
	viper.AutomaticEnv()

	if vipFile == "" {
		if viper.GetBool("debug") {
			log.Println("Configuration file ignored.")
		}
		return
	}

	viper.SetConfigFile(vipFile)
	viper.SafeWriteConfigAs(vipFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

}

func initDataset() {

	datadir := viper.GetString("dataset.dir")

	if datadir != "" {
		os.MkdirAll(datadir, 0755)
	} else {
		viper.Set("dataset.dir", ".")
	}

}

func initLogger() {

	logdir := viper.GetString("logger.dir")

	if logdir != "" {
		os.MkdirAll(logdir, 0755)
	} else {
		viper.Set("logger.dir", ".")
	}

	if !viper.GetBool("logger.output") {
		return
	}

	file := logdir + "/output.log"
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	if logfile, err := os.OpenFile(file, flag, 0644); err == nil {
		log.SetOutput(io.MultiWriter(os.Stdout, logfile))
	} else {
		log.Println("Failed to", err)
	}

}
