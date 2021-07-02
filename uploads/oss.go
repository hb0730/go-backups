package uploads

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	util "github.com/hb0730/go-backups/utils"
	"os"
	"time"
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
	newFilename, err := ali.getTempFile(filename)
	if err != nil {
		return err
	}
	err = compress.CompressDir(newFilename, dir)
	if err != nil {
		return err
	}
	return ali.bucket.PutObjectFromFile(ali.getUploadFilenameWithPath(filename), newFilename)
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
	//得到的是临时文件(存在文件路径)
	localFile, err := ali.getTempFile(filename)
	if err != nil {
		return err
	}
	err = compress.CompressDirs(localFile, dirs)
	if err != nil {
		return err
	}
	return ali.after(util.GetFilenameWithExtension(localFile), localFile)

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
func (ali *AliyunOss) after(filename, localFile string) error {
	return ali.bucket.PutObjectFromFile(ali.getUploadFilenameWithPath(filename), localFile)
}

func (ali *AliyunOss) validate() error {
	if ali.Endpoint == "" || ali.AccessKeyId == "" || ali.AccessKeySecret == "" || ali.BucketName == "" {
		return errors.New("field required")
	}
	return nil
}
func (ali *AliyunOss) getTempFile(filename string) (string, error) {
	tempDir, err := util.GetTempDir("")
	if err != nil {
		return "", err
	}
	return tempDir + string(os.PathSeparator) + filename, nil
}
func (ali *AliyunOss) getUploadFilenameWithPath(filename string) string {
	return time.Now().Format("2006-01-02") + "/" + filename
}
