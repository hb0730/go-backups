package uploads

import (
	"testing"
)

const (
	endpoint        = ""
	bucketName      = ""
	accessKeyId     = ""
	accessKeySecret = ""
)

func TestAliyunOss_UploadDir(t *testing.T) {
	ali, err := NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName)
	if err != nil {
		panic(err)
	}
	err = ali.UploadDir(`E:\goWork\zinx-demo`, `zinx-demo.zip`, "")
	if err != nil {
		panic(err)
	}
}
