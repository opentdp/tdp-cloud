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

	args.Parser()

	switch args.CmdName {
	case "server":
		server.Create(&vfs)
	case "worker":
		worker.Create()
	}

}
