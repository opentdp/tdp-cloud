package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/global"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/version"
	"tdp-cloud/cmd/worker"
)

func Execute() {

	cobra.OnInitialize(initViper)

	cli := global.WithCli()

	cli.AddCommand(
		server.WithCli(), worker.WithCli(), version.WithCli(),
	)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}

}

func initViper() {

	viper.AutomaticEnv()

	if args.ConfigFile == "" {
		log.Println("Config file will be ignored")
		return
	}

	viper.SetConfigFile(args.ConfigFile)
	viper.SafeWriteConfigAs(args.ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}

}
