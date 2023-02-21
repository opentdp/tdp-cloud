package cmd

import (
	"os"

	"tdp-cloud/cmd/subset"
)

func Execute() {

	rootCmd.AddCommand(
		subset.WithServer(), subset.WithWorker(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
