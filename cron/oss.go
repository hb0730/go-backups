package cron

import (
	"github.com/hb0730/go-backups/config"
	"github.com/hb0730/go-backups/uploads"
)

type AliYunOss struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	BucketName      string `json:"bucket_name"`
}

func (ali *AliYunOss) Upload(name, supportKey string) error {
	err := config.ReadYaml().Unmarshal(name+"."+supportKey, ali)
	if err != nil {
		return err
	}

	return nil
}
func (ali *AliYunOss) UploadFromAliYunOss(name string) error {
	oss, err := uploads.NewAliyunOss(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret, ali.BucketName)
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
