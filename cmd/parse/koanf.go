package parse

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"

	"tdp-cloud/cmd/args"
)

// 配置文件路径

var YamlFile string

// 配置信息操作类

type Config struct {
	Koanf    *koanf.Koanf
	Parser   *yaml.YAML
	Override bool
}

func (c *Config) Init() *Config {

	debug := os.Getenv("TDP_DEBUG")
	args.Debug = debug == "1" || debug == "true"

	c.Parser = yaml.Parser()
	c.Koanf = koanf.NewWithConf(koanf.Conf{
		StrictMerge: true,
		Delim:       ".",
	})

	return c

}

func (c *Config) ReadYaml() {

	// 不存在则忽略
	if _, err := os.Stat(YamlFile); os.IsNotExist(err) {
		return
	}

	// 从文件读取参数
	err := c.Koanf.Load(file.Provider(YamlFile), c.Parser)
	if err != nil {
		logman.Fatal("read config failed", "error", err)
	}

}

func (c *Config) WriteYaml() {

	// 是否强制覆盖
	if !c.Override && filer.Exists(YamlFile) {
		return
	}

	// 序列化参数信息
	bytes, err := c.Koanf.Marshal(c.Parser)
	if err != nil {
		logman.Fatal("write config failed", "error", err)
	}

	// 将参数写入配置文件
	err = os.WriteFile(YamlFile, bytes, 0644)
	if err != nil {
		logman.Fatal("write config failed", "error", err)
	}

}
