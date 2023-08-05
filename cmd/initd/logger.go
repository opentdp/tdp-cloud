package initd

import (
	"os"

	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

func Logger() {

	logdir := args.Logger.Dir

	if logdir != "" && logdir != "." {
		os.MkdirAll(logdir, 0755)
	}

	logman.SetDefault(&logman.Config{
		Level:    args.Logger.Level,
		Target:   args.Logger.Target,
		Storage:  logdir,
		Filename: "global",
	})

}
