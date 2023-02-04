package version

import (
	"fmt"

	"tdp-cloud/cmd/args"
)

func Execute() {

	fmt.Println("Version", args.Version)
	fmt.Println("Build", args.BuildVersion)

}
