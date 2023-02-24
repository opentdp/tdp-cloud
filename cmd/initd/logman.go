package initd

import (
	"os"

	"github.com/spf13/viper"

	"tdp-cloud/helper/logman"
)

func Logman() {

	logdir := viper.GetString("logger.dir")

	if logdir != "" {
		os.MkdirAll(logdir, 0755)
	} else {
		viper.Set("logger.dir", ".")
	}

	logman.New()

}
