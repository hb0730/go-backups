package cron

import (
	"fmt"
	"github.com/hb0730/go-backups/config"
	"github.com/hb0730/go-backups/uploads"
	"github.com/knadh/koanf"
	"github.com/mitchellh/mapstructure"
	"github.com/mritd/logger"
)

var Supports = make(map[string]Support)

type Support interface {
	//Upload
	// name 上传的名称: blog
	// supportType 上传的类型: aliyun-oss
	// 最后得到的则是blog.aliyun-oss
	Upload(name, supportType string) error
}

type Base struct {
	//Compress compress type
	//see utils.CompressSuports
	Compress string `json:"compress"`
	upload   uploads.Upload
}

func (b *Base) Unmarshal(name, supportKey string, g Support) error {
	logger.Infof("read config key: %s.%s", name, supportKey)
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

//Uploads 上传
// name 上传的名称
func (b *Base) Uploads(name string) error {
	filename, dirs := readFileYaml(name)
	return b.upload.UploadDirs(dirs, filename, "")
}
func readFileYaml(name string) (filename string, dirs []string) {
	filename = config.ReadYaml().String(name + ".uploads.filename")
	dirs = config.ReadYaml().Strings(name + ".uploads.dirs")
	return
}
