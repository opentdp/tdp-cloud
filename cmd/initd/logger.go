package initd

import (
	"os"

	"tdp-cloud/cmd/args"
	"tdp-cloud/helper/logman"
)

func Logger() {

	logdir := args.Logger.Dir

	if logdir != "" && logdir != "." {
		os.MkdirAll(logdir, 0755)
	}

	logman.New()

}
