package cron

import (
	"fmt"
	"github.com/hb0730/go-backups/config"
	"github.com/knadh/koanf"
	"github.com/mitchellh/mapstructure"
)

var Supports = make(map[string]Support)

type Support interface {
	//Upload
	Upload(name, supportType string) error
}

type Base struct {
	//Compress compress type
	//see utils.CompressSuports
	Compress string `json:"compress"`
}

func (b *Base) Unmarshal(name, supportKey string, g Support) error {
	return config.ReadYaml().UnmarshalWithConf(
		fmt.Sprintf("%s.%s", name, supportKey),
		g,
		koanf.UnmarshalConf{
			DecoderConfig: &mapstructure.DecoderConfig{
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc()),
				Metadata:         nil,
				Result:           g,
				WeaklyTypedInput: true,
				Squash:           true,
			},
		})
}
