package subset

import (
	"flag"

	"tdp-cloud/cmd/parse"
	"tdp-cloud/service"
)

var workerYaml string

func workerFlag() *FlagSet {

	var action string

	command := &FlagSet{
		FlagSet: flag.NewFlagSet("worker", flag.ExitOnError),
		Comment: "TDP Cloud Worker Management",
		Execute: func() {
			workerExec(action)
		},
	}

	command.StringVar(&action, "s", "", "management worker service")
	command.StringVar(&workerYaml, "c", "", "config file path")

	return command

}

func workerExec(act string) {

	c := parse.WorkerConfig(workerYaml)

	if act == "" || act == "start" {
		c.Save()
	}

	service.Control("worker", act)

}
