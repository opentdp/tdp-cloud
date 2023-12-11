package parse

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
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
	Koanf  *koanf.Koanf
	Parser *yaml.YAML
}

func NewConfig() *Config {

	var p = yaml.Parser()
	var k = koanf.NewWithConf(koanf.Conf{
		StrictMerge: true,
		Delim:       ".",
	})

	return &Config{k, p}

}

func (c *Config) Server() {

	// 读取默认配置
	mp := map[string]any{
		"dataset":  args.Dataset,
		"logger":   args.Logger,
		"database": args.Database,
		"server":   args.Server,
	}
	c.Koanf.Load(confmap.Provider(mp, "."), nil)

	// 读取配置文件
	if YamlFile != "" {
		c.ReadYaml()
		// 使用配置文件中的参数覆盖默认值
		c.Koanf.Unmarshal("dataset", &args.Dataset)
		c.Koanf.Unmarshal("logger", &args.Logger)
		c.Koanf.Unmarshal("database", &args.Database)
		c.Koanf.Unmarshal("server", &args.Server)
	}

	RuntimeFix()

}

func (c *Config) Worker() {

	// 读取默认配置
	mp := map[string]any{
		"dataset": args.Dataset,
		"logger":  args.Logger,
		"worker":  args.Worker,
	}
	c.Koanf.Load(confmap.Provider(mp, "."), nil)

	// 读取配置文件
	if YamlFile != "" {
		c.ReadYaml()
		// 使用配置文件中的参数覆盖默认值
		c.Koanf.Unmarshal("dataset", &args.Dataset)
		c.Koanf.Unmarshal("logger", &args.Logger)
		c.Koanf.Unmarshal("worker", &args.Worker)
	}

	RuntimeFix()

}

func (c *Config) ReadYaml() {

	// 配置不存在则忽略
	_, err := os.Stat(YamlFile)
	if os.IsNotExist(err) {
		return
	}

	// 从配置文件读取参数
	err = c.Koanf.Load(file.Provider(YamlFile), c.Parser)
	if err != nil {
		logman.Fatal("read config error", "error", err)
	}

}

func (c *Config) WriteYaml(force bool) {

	// 是否强制覆盖
	if !force && filer.Exists(YamlFile) {
		return
	}

	// 序列化参数信息
	b, err := c.Koanf.Marshal(c.Parser)
	if err != nil {
		logman.Fatal("write config error", "error", err)
	}

	// 将参数写入配置文件
	err = os.WriteFile(YamlFile, b, 0644)
	if err != nil {
		logman.Fatal("write config error", "error", err)
	}

}
