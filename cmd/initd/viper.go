package initd

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
)

var ViperFile = ""

func Viper() {

	defer args.Load()

	// 环境变量

	viper.SetEnvPrefix("TDP")
	viper.AutomaticEnv()

	// 配置文件

	if ViperFile == "" {
		log.Fatal("Config file must be specified")
	}

	viper.SetConfigFile(ViperFile)

	// 读取配置

	if _, err := os.Stat(ViperFile); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

}
