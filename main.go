package main

import (
	"embed"
	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/worker"
)

//go:embed front
var vfs embed.FS

func main() {

	args.Flags()

	if args.Server == "" {
		server.Create(&vfs)
	} else {
		worker.Create()
	}

}
