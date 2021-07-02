package uploads

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	util "github.com/hb0730/go-backups/utils"
	"os"
	"time"
)

type AliyunOss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	bucket          *oss.Bucket
}

func NewAliyunOss(endpoint, accessKeyId, accessKeySecret, bucketName string) (*AliyunOss, error) {
	ali := new(AliyunOss)
	ali.Endpoint = endpoint
	ali.AccessKeyId = accessKeyId
	ali.AccessKeySecret = accessKeySecret
	ali.BucketName = bucketName
	return ali, ali.Client()
}

func (ali *AliyunOss) Client() (err error) {
	aliClient, err := oss.New(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret)
	if err != nil {
		return err
	}
	ali.bucket, err = aliClient.Bucket(ali.BucketName)
	return
}

func (ali *AliyunOss) UploadDir(dir, filename, _ string) error {
	err := ali.before()
	if err != nil {
		return err
	}
	z := util.NewZipUtils()
	defer z.Close()
	newFilename, err := ali.getTempFile(filename)
	if err != nil {
		return err
	}
	err = z.CompressDir(dir, newFilename)
	if err != nil {
		return err
	}
	return ali.bucket.PutObjectFromFile(ali.getUploadFilenameWithPath(filename), newFilename)
}

func (ali *AliyunOss) UploadDirs(dirs []string, filename, _ string) error {
	err := ali.before()
	if err != nil {
		return err
	}
	z := util.NewZipUtils()
	defer z.Close()
	//得到的是临时文件(存在文件路径)
	localFile, err := ali.getTempFile(filename)
	if err != nil {
		return err
	}
	err = z.CompressDirs(dirs, localFile)
	if err != nil {
		return err
	}
	return ali.after(util.GetFilenameWithExtension(localFile), localFile)

}
func (ali *AliyunOss) before() error {
	err := ali.validate()
	if err != nil {
		return err
	}
	if ali.bucket == nil {
		err = ali.Client()
		if err != nil {
			return err
		}
	}
	return nil
}
func (ali AliyunOss) after(filename, localFile string) error {
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
func (ali AliyunOss) getUploadFilenameWithPath(filename string) string {
	return time.Now().Format("2006-01-02") + "/" + filename
}
