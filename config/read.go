package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k *koanf.Koanf

func ReadYaml() *koanf.Koanf {
	if k == nil {
		return LoadKoanf("")
	}
	return k
}
func LoadKoanf(config string) *koanf.Koanf {
	load(config)
	return k
}
func load(config string) {
	k = koanf.New(".")
	if config == "" {
		defaultConfig()
	} else {
		err := k.Load(file.Provider(config), yaml.Parser())
		if err != nil {
			panic(err)
		}
	}
}

func defaultConfig() {
	_ = k.Load(file.Provider("./config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("../config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("/config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("application.yml"), yaml.Parser())
}
