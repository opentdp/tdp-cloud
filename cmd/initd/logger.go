package initd

import (
	"os"

	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
)

func init() {

	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.tofile", false)
	viper.SetDefault("logger.stdout", true)

}

func Logger() {

	logdir := viper.GetString("logger.dir")

	if logdir != "" {
		os.MkdirAll(logdir, 0755)
	} else {
		viper.Set("logger.dir", ".")
	}

	logman.New()

}
