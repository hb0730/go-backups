package uploads

import (
	"context"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	util "github.com/hb0730/go-backups/utils"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type AliyunOss struct {
	BaseUpload
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	bucket          *oss.Bucket
}

func NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName, compress string) (*AliyunOss, error) {
	ali := new(AliyunOss)
	ali.Endpoint = endpoint
	ali.AccessKeyId = accessKeyId
	ali.AccessKeySecret = accessKeySecret
	ali.BucketName = bucketName
	ali.CompressType = compress
	return ali, ali.Client()
}
func (ali *AliyunOss) SetBucket(bucket *oss.Bucket) *AliyunOss {
	ali.bucket = bucket
	return ali
}

func (ali *AliyunOss) Client() (err error) {
	err = ali.validate()
	if err != nil {
		return err
	}
	aliClient, err := oss.New(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret)
	if err != nil {
		return err
	}
	ali.bucket, err = aliClient.Bucket(ali.BucketName)
	return
}

func (ali *AliyunOss) UploadDir(dir, filename, _ string) error {
	if dir == "" || filename == "" {
		return nil
	}
	compress, err := ali.before()
	if err != nil {
		return err
	}

	defer compress.Close()
	//获取临时目录 temp/filename
	newFilepath, err := util.GetTempDirFile("", filename)
	if err != nil {
		return err
	}
	// temp/filename.extension
	err = compress.CompressDir(&newFilepath, dir)
	if err != nil {
		return err
	}
	newFilename := util.UploadFilenameWithDir(newFilepath)
	return ali.bucket.PutObjectFromFile(newFilename, newFilepath)
}

func (ali *AliyunOss) UploadDirs(dirs []string, filename, _ string) error {
	if len(dirs) == 0 || filename == "" {
		return nil
	}
	compress, err := ali.before()
	if err != nil {
		return err
	}
	defer compress.Close()
	newFilePath, err := util.GetTempDirFile("", filename)
	if err != nil {
		return err
	}
	err = compress.CompressDirs(&newFilePath, dirs)
	if err != nil {
		return err
	}
	newFilename := util.UploadFilenameWithDir(newFilePath)
	return ali.bucket.PutObjectFromFile(newFilename, newFilePath)

}
func (ali *AliyunOss) before() (util.Compress, error) {
	err := ali.validate()
	if err != nil {
		return nil, err
	}
	if ali.bucket == nil {
		err = ali.Client()
		if err != nil {
			return nil, err
		}
	}
	return ali.GetCompress()
}

func (ali *AliyunOss) validate() error {
	if ali.Endpoint == "" || ali.AccessKeyId == "" || ali.AccessKeySecret == "" || ali.BucketName == "" {
		return errors.New("field required")
	}
	return nil
}

type QiniuOss struct {
	BaseUpload
	AccessKey    string
	SecretKey    string
	Bucket       string
	RegionId     string
	upToken      string
	fromUploader *storage.FormUploader
}

func NewQiniuOss(accessKey, secretKey, bucket, regionId, compressType string) (*QiniuOss, error) {
	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	re, is := storage.GetRegionByID(storage.RegionID(regionId))
	if !is {
		return nil, errors.New("region not found ")
	}
	cfg.Zone = &re
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	// 构建表单上传的对象
	fromUploader := storage.NewFormUploader(&cfg)
	q := new(QiniuOss)
	q.AccessKey = accessKey
	q.SecretKey = secretKey
	q.Bucket = bucket
	q.RegionId = regionId
	q.CompressType = compressType
	q.upToken = upToken
	q.fromUploader = fromUploader
	return q, nil
}
func (q *QiniuOss) SetUpToken(upToken string) *QiniuOss {
	q.upToken = upToken
	return q
}
func (q *QiniuOss) SetFromUploader(uploader *storage.FormUploader) *QiniuOss {
	q.fromUploader = uploader
	return q
}

func (q *QiniuOss) UploadDir(dir, filename string, _ string) error {
	if dir == "" {
		return errors.New("Directory is not blank")
	}
	if filename == "" {
		return errors.New("filename is not blank")
	}
	err := q.validate()
	if err != nil {
		return err
	}
	compress, err := q.GetCompress()
	if err != nil {
		return err
	}
	defer compress.Close()
	newFilepath, err := util.GetTempDirFile("", filename)
	if err != nil {
		return err
	}
	err = compress.CompressDir(&newFilepath, dir)
	if err != nil {
		return err
	}
	newFilename := util.UploadFilenameWithDir(newFilepath)
	return q.fromUploader.PutFile(context.Background(), nil, q.upToken, newFilename, newFilepath, nil)
}

func (q *QiniuOss) UploadDirs(dir []string, filename string, _ string) error {
	if len(dir) == 0 {
		return errors.New("Directory is not empty")
	}
	if filename == "" {
		return errors.New("filename is not blank")
	}
	err := q.validate()
	if err != nil {
		return err
	}
	compress, err := q.GetCompress()
	if err != nil {
		return err
	}
	defer compress.Close()
	newFilepath, err := util.GetTempDirFile("", filename)
	if err != nil {
		return err
	}
	err = compress.CompressDirs(&newFilepath, dir)
	if err != nil {
		return err
	}
	newFilename := util.UploadFilenameWithDir(newFilepath)
	return q.fromUploader.PutFile(context.Background(), nil, q.upToken, newFilename, newFilepath, nil)
}

func (q *QiniuOss) validate() error {
	if q.upToken == "" || q.fromUploader == nil {
		return errors.New("uninitialized")
	}
	return nil
}
