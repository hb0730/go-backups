package cron

import (
	"github.com/hb0730/go-backups/config"
	"testing"
)

func TestAliYunOss_Upload(t *testing.T) {
	config.LoadKoanf("")
	ali := AliYunOss{}
	err := ali.Upload("blog", "aliyun-oss")
	if err != nil {
		panic(err)
	}
}
