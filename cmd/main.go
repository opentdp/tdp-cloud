package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/global"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/service"
	"tdp-cloud/cmd/version"
	"tdp-cloud/cmd/worker"
)

func Execute() {

	cobra.OnInitialize(initViper)

	cli := global.WithCli()

	cli.AddCommand(
		server.WithCli(),
		worker.WithCli(),
		service.WithCli(),
		version.WithCli(),
	)

	if err := cli.Execute(); err != nil {
		panic(err)
	}

}

func initViper() {

	viper.AutomaticEnv()

	viper.SetConfigFile(args.ConfigFile)
	viper.SafeWriteConfigAs(args.ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
