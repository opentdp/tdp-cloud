package initd

import (
	"os"

	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
)

func init() {

	viper.SetDefault("logger.dir", ".")
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.stdout", true)
	viper.SetDefault("logger.tofile", false)

}

func Logger() {

	logdir := viper.GetString("logger.dir")

	if logdir != "" && logdir != "." {
		os.MkdirAll(logdir, 0755)
	}

	logman.New()

}
