package parse

import (
	"os"
	"path"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

func preConfig() {

	debug := os.Getenv("TDP_DEBUG")
	args.Debug = debug == "1" || debug == "true"

}

func postConfig() {

	// init dataset

	if args.Dataset.Secret == "" {
		args.Dataset.Secret = strutil.Rand(32)
	}

	if args.Dataset.Dir != "" && args.Dataset.Dir != "." {
		os.MkdirAll(args.Dataset.Dir, 0755)
	}

	// init logger

	if !path.IsAbs(args.Logger.Dir) {
		args.Logger.Dir = path.Join(args.Dataset.Dir, args.Logger.Dir)
	}

	if args.Logger.Dir != "" && args.Logger.Dir != "." {
		os.MkdirAll(args.Logger.Dir, 0755)
	}

	logman.SetDefault(&logman.Config{
		Level:    args.Logger.Level,
		Target:   args.Logger.Target,
		Storage:  args.Logger.Dir,
		Filename: "default",
	})

}
