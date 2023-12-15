package parse

import (
	"errors"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
)

// 配置文件路径

var YamlFile string

// 配置信息操作类

type Config struct {
	Koanf    *koanf.Koanf
	Parser   *yaml.YAML
	Override bool
}

func NewConfig() *Config {

	return &Config{
		Parser: yaml.Parser(),
		Koanf: koanf.NewWithConf(koanf.Conf{
			StrictMerge: true,
			Delim:       ".",
		}),
	}

}

func (c *Config) ReadYaml() error {

	if YamlFile == "" {
		return errors.New("config file not set")
	}

	// 不存在则忽略
	if _, err := os.Stat(YamlFile); os.IsNotExist(err) {
		return err
	}

	// 从文件读取参数
	err := c.Koanf.Load(file.Provider(YamlFile), c.Parser)
	if err != nil {
		logman.Fatal("read config failed", "error", err)
	}

	return err

}

func (c *Config) WriteYaml() error {

	if YamlFile == "" {
		return errors.New("config file not set")
	}

	// 是否强制覆盖
	if !c.Override && filer.Exists(YamlFile) {
		return errors.New("config file not exist")
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

	return err

}
