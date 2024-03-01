package parse

import (
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

type WorkerData struct {
	Assets *args.IAssets `yaml:"dataset"`
	Logger *args.ILogger `yaml:"logger"`
	Worker *args.IWorker `yaml:"worker"`
}

func WorkerConfig(yaml string) *Config {

	args.SetDebug()

	config := &Config{
		File: yaml,
		Data: &WorkerData{
			args.Assets, args.Logger, args.Worker,
		},
	}

	if err := config.Load(); err != nil {
		logman.Fatal("load config failed", "error", err)
	}

	args.SetAssets()
	args.SetLogger()

	return config

}
