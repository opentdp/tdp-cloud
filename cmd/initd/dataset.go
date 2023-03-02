package initd

import (
	"os"

	"tdp-cloud/cmd/args"
)

func Dataset() {

	datadir := args.Dataset.Dir

	if datadir != "" && datadir != "." {
		os.MkdirAll(datadir, 0755)
	}

}
