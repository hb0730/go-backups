package cron

import (
	"github.com/hb0730/go-backups/config"
	"github.com/hb0730/go-backups/uploads"
)

type AliYunOss struct {
	Base            `mapstructure:",squash"`
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
}

func (ali *AliYunOss) Upload(name, supportKey string) error {
	err := ali.Unmarshal(name, supportKey, ali)
	if err != nil {
		return err
	}
	return ali.UploadFromAliYunOss(name)
}
func (ali *AliYunOss) UploadFromAliYunOss(name string) error {
	oss, err := uploads.NewAliyunOss(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret, ali.BucketName, ali.Compress)
	if err != nil {
		return err
	}
	filename := config.ReadYaml().String(name + ".uploads.filename")
	dirs := config.ReadYaml().Strings(name + ".uploads.dirs")
	return oss.UploadDirs(dirs, filename, "")
}

func init() {
	Supports["aliyun-oss"] = &AliYunOss{}
}
