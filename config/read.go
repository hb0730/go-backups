package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k *koanf.Koanf

func ReadYaml() *koanf.Koanf {
	if k == nil {
		return LoadKoanf()
	}
	return k
}
func LoadKoanf() *koanf.Koanf {
	load()
	return k
}
func load() {
	k = koanf.New(".")
	//err := k.Load(file.Provider("./config/application.yml"), yaml.Parser())
	//err = k.Load(file.Provider("../config/application.yml"), yaml.Parser())
	//err = k.Load(file.Provider("/config/application.yml"), yaml.Parser())
	//err = k.Load(file.Provider("config/application.yml"), yaml.Parser())
	//err = k.Load(file.Provider("application.yml"), yaml.Parser())
	//if err != nil {
	//	panic(err)
	//}
	_ = k.Load(file.Provider("./config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("../config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("/config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("application.yml"), yaml.Parser())

}
