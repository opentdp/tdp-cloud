package main

import (
	"embed"

	"tdp-cloud/cmd"
	"tdp-cloud/cmd/args"
)

//go:embed front
var vfs embed.FS

func main() {

	args.FrontFS = &vfs

	cmd.Execute()

}
