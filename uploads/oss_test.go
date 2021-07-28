package uploads

import (
	"testing"
)

const (
	endpoint        = ""
	bucketName      = "vback"
	accessKeyId     = "1HAzkkzUqhj2XZQ5-ivsk_Vo53Tw-bHz5WikULVL"
	accessKeySecret = "ie_U79nV07VzTZuDKq-cIphPzyNGpYmfWKD5_k37"
	regionId        = "z0"
)

func TestAliyunOss_UploadDir(t *testing.T) {
	ali, err := NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName, "zip")
	if err != nil {
		panic(err)
	}
	err = ali.UploadDir(`E:\goWork\zinx-demo`, `zinx-demo`, "")
	if err != nil {
		panic(err)
	}
}
func TestAliyunOss_UploadDirs(t *testing.T) {
	ali, err := NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName, "zip")
	if err != nil {
		panic(err)
	}
	err = ali.UploadDirs([]string{`E:\goWork\zinx-demo`, `E:\goWork\subdomain`}, `demo`, "")
	if err != nil {
		panic(err)
	}
}
func TestQiniuOss_UploadDir(t *testing.T) {
	q, err := NewQiniuOss(accessKeyId, accessKeySecret, bucketName, regionId, "zip")
	if err != nil {
		panic(err)
	}
	err = q.UploadDir(`E:\goWork\zinx-demo`, `zinx-demo`, "")
	if err != nil {
		panic(err)
	}
}

func TestQiniuOss_UploadDirs(t *testing.T) {
	q, err := NewQiniuOss(accessKeyId, accessKeySecret, bucketName, regionId, "zip")
	if err != nil {
		panic(err)
	}
	err = q.UploadDirs([]string{`E:\goWork\zinx-demo`, `E:\goWork\subdomain`}, `demo`, "")
	if err != nil {
		panic(err)
	}
}
