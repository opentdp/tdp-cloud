package subset

import (
	"flag"

	"tdp-cloud/cmd/parse"
	"tdp-cloud/service"
)

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
	command.StringVar(&parse.YamlFile, "c", "config.yml", "config file path")

	return command

}

func workerExec(act string) {

	c := parse.NewConfig()
	c.Worker()

	if act == "" || act == "start" {
		c.WriteYaml(false)
	}

	service.Control("worker", act)

}
