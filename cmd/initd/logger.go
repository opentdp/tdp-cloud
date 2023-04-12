package initd

import (
	"os"

	"github.com/open-tdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

func Logger() {

	logdir := args.Logger.Dir

	if logdir != "" && logdir != "." {
		os.MkdirAll(logdir, 0755)
	}

	logman.SetDefault(&logman.Param{
		Level:    args.Logger.Level,
		Target:   args.Logger.Target,
		Filename: logdir + "/" + args.Logger.Level + ".log",
	})

}
