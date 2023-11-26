package parse

import (
	"os"
	"path"
	"tdp-cloud/cmd/args"

	"github.com/opentdp/go-helper/logman"
)

func RuntimeFix() {

	// init dataset

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
		Filename: "global",
	})

}
