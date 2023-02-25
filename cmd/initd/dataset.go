package initd

import (
	"os"

	"github.com/spf13/viper"
)

func init() {

	viper.SetDefault("logger.dir", ".")

}

func Dataset() {

	datadir := viper.GetString("dataset.dir")

	if datadir != "" && datadir != "." {
		os.MkdirAll(datadir, 0755)
	}

}
