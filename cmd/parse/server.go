package parse

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

type ServerData struct {
	Assets *args.IAssets `yaml:"dataset"`
	Gormio *args.IGormio `yaml:"database"`
	Logger *args.ILogger `yaml:"logger"`
	Server *args.IServer `yaml:"server"`
}

func ServerConfig(yaml string) *Config {

	args.SetDebug()

	config := &Config{
		File: yaml,
		Data: &ServerData{
			args.Assets, args.Gormio, args.Logger, args.Server,
		},
	}

	if config.File == "" {
		config.setYaml("server")
	}

	if err := config.Load(); err != nil {
		logman.Fatal("load config failed", "error", err)
	}

	if args.Gormio.Type == "sqlite" {
		args.Gormio.Host = args.Assets.Dir
	}

	if args.Server.JwtKey == "" {
		args.Server.JwtKey = strutil.Rand(32)
	}

	args.SetAssets()
	args.SetLogger()

	return config

}
