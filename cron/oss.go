package cron

import (
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
	ali.upload = oss
	return ali.Uploads(name)
}

type QiniuOss struct {
	Base      `mapstructure:",squash"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	RegionId  string `json:"regionId"`
}

func (q *QiniuOss) Upload(name, supportKey string) error {
	err := q.Unmarshal(name, supportKey, q)
	if err != nil {
		return err
	}
	oss, err := uploads.NewQiniuOss(q.AccessKey, q.SecretKey, q.Bucket, q.RegionId, q.Compress)
	if err != nil {
		return err
	}
	q.upload = oss
	return q.Uploads(name)
}

func init() {
	Supports["aliyun-oss"] = &AliYunOss{}
	Supports["qiniu-oss"] = &QiniuOss{}
}
