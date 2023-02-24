package initd

import (
	"os"

	"github.com/spf13/viper"
)

func Dataset() {

	datadir := viper.GetString("dataset.dir")

	if datadir != "" {
		os.MkdirAll(datadir, 0755)
	} else {
		viper.Set("dataset.dir", ".")
	}

}
