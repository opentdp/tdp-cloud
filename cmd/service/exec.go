package service

import (
	"github.com/spf13/viper"

	"tdp-cloud/module/installer"
)

func Execute() {

	in := viper.GetString("install")
	un := viper.GetString("uninstall")

	installer.Execute(in, un)

}
