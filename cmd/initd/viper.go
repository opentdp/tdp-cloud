package initd

import (
	"log"
	"tdp-cloud/cmd/args"

	"github.com/spf13/viper"
)

var ViperFile = ""

func Viper() {

	defer args.Sync()

	if ViperFile == "" {
		log.Fatal("Configuration file must be specified")
	}

	// 环境变量

	viper.SetEnvPrefix("TDP")
	viper.AutomaticEnv()

	// 安全写入

	viper.SetConfigFile(ViperFile)
	viper.SafeWriteConfigAs(ViperFile)

	// 读取配置

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// 强制更新

	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}

}
