package initd

import (
	"log"

	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
)

func Viper() {

	// 环境变量

	viper.SetEnvPrefix("TDP")
	viper.AutomaticEnv()

	// 忽略配置

	if args.ConfigFile == "" {
		return
	}

	// 写入配置

	viper.SetConfigFile(args.ConfigFile)
	viper.SafeWriteConfigAs(args.ConfigFile)

	// 读取配置

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}
