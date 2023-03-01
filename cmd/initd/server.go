package initd

import (
	"github.com/spf13/viper"
)

func init() {

	viper.SetDefault("server.register", true)

}
