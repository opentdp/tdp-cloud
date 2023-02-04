package main

import (
	"embed"

	"tdp-cloud/cmd"
	"tdp-cloud/cmd/args"
)

//go:embed front
var efs embed.FS

func main() {

	args.EmbedFs = &efs

	cmd.Execute()

}
