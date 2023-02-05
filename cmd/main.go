package cmd

import (
	"os"

	"tdp-cloud/cmd/subset"
)

func Execute() {

	rcmd.AddCommand(
		subset.WithServer(), subset.WithWorker(), subset.WithVersion(),
	)

	if err := rcmd.Execute(); err != nil {
		os.Exit(1)
	}

}
