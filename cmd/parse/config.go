package parse

import (
	"os"

	"github.com/opentdp/go-helper/filer"
	"gopkg.in/yaml.v3"
)

type Config struct {
	File string
	Data any
}

func (c *Config) Load() error {

	if c.File == "" {
		return nil
	}

	if !filer.Exists(c.File) {
		return nil
	}

	bytes, err := os.ReadFile(c.File)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, c.Data)

}

func (c *Config) Save() error {

	if c.File == "" {
		return nil
	}

	bytes, err := yaml.Marshal(c.Data)
	if err != nil {
		return err
	}

	return os.WriteFile(c.File, bytes, 0644)

}

func (c *Config) setYaml(n string) {

	if f := "config.yml"; filer.Exists(f) {
		c.File = f
		return
	}

	if f := "/etc/tdp-" + n + ".yml"; filer.Exists(f) {
		c.File = f
		return
	}

	if f := "/etc/tdp-cloud/" + n + ".yml"; filer.Exists(f) {
		c.File = f
		return
	}

}
