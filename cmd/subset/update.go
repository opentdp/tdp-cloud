package subset

import (
	"flag"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/upgrade"

	"tdp-cloud/cmd/args"
)

func updateFlag() *FlagSet {

	command := &FlagSet{
		FlagSet: flag.NewFlagSet("update", flag.ExitOnError),
		Comment: "TDP Cloud Update Management",
		Execute: func() {
			updateExec()
		},
	}

	return command

}

func updateExec() {

	err := upgrade.Apply(&upgrade.RequesParam{
		Server:  args.UpdateUrl,
		Version: args.Version,
	})

	if err == nil {
		logman.Info("Update Success")
	}

}
