package main

import (
	"embed"
	"os"

	"tdp-cloud/cmd"
	"tdp-cloud/cmd/args"
	"tdp-cloud/cmd/server"
	"tdp-cloud/cmd/service"
	"tdp-cloud/cmd/worker"
)

//go:embed front
var vfs embed.FS

func main() {

	args.Parser()
	cmd.FrontFS = &vfs

	switch os.Args[1] {
	case "server":
		server.Create()
	case "service":
		service.Create()
	case "worker":
		worker.Create()
	}

}
