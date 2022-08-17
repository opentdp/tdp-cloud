package main

import (
	"embed"
	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/master"
	"tdp-cloud/cmd/worker"
)

//go:embed front
var vfs embed.FS

func main() {

	args.Flags()

	if args.Master == "" {
		master.Create(&vfs)
	} else {
		worker.Create()
	}

}
