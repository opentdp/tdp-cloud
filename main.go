package main

import (
	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/master"
	"tdp-cloud/cmd/worker"
)

func main() {

	args.Flags()

	if args.Master == "" {
		master.Create(args.Listen)
	} else {
		worker.Create(args.Master)
	}

}
