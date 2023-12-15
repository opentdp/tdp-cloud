package parse

import (
	"os"
	"path"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"

	"tdp-cloud/cmd/args"
)

func (c *Config) Server() {

	// 读取默认配置

	mp := map[string]any{
		"dataset":  &args.Dataset,
		"database": &args.Database,
		"logger":   &args.Logger,
		"server":   &args.Server,
	}
	c.Koanf.Load(confmap.Provider(mp, "."), nil)

	// 读取配置文件

	if YamlFile != "" {
		c.ReadYaml()
		for k, v := range mp {
			c.Koanf.Unmarshal(k, v)
		}
	}

	// 初始化存储目录

	if args.Dataset.Secret == "" {
		args.Dataset.Secret = strutil.Rand(32)
		c.Override = true
	}

	if args.Dataset.Dir != "" && args.Dataset.Dir != "." {
		os.MkdirAll(args.Dataset.Dir, 0755)
	}

	// 修正数据库参数

	if args.Database.Type == "sqlite" && !path.IsAbs(args.Database.Name) {
		args.Database.Name = path.Join(args.Dataset.Dir, args.Database.Name)
	}

	// 初始化日志能力

	if !path.IsAbs(args.Logger.Dir) {
		args.Logger.Dir = path.Join(args.Dataset.Dir, args.Logger.Dir)
	}

	if args.Logger.Dir != "" && args.Logger.Dir != "." {
		os.MkdirAll(args.Logger.Dir, 0755)
	}

	logman.SetDefault(&logman.Config{
		Level:    args.Logger.Level,
		Target:   args.Logger.Target,
		Storage:  args.Logger.Dir,
		Filename: "server",
	})

	// 初始化 JwtKey

	if args.Server.JwtKey == "" {
		args.Server.JwtKey = strutil.Rand(32)
		c.Override = true
	}

}
