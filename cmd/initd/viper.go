package initd

import (
	"log"
	"tdp-cloud/cmd/args"

	"github.com/spf13/viper"
)

var ViperFile = ""

func Viper() {

	defer args.Sync()

	// 环境变量

	viper.SetEnvPrefix("TDP")
	viper.AutomaticEnv()

	// 忽略配置

	if ViperFile == "" {
		return
	}

	// 写入配置

	viper.SetConfigFile(ViperFile)
	viper.SafeWriteConfigAs(ViperFile)

	// 读取配置

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}
