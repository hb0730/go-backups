package uploads

import (
	"errors"
	util "github.com/hb0730/go-backups/utils"
)

type Upload interface {
	//UploadDir git upload compress file
	//dir compress dir
	//filename File with name only
	//description description
	UploadDir(dir, filename string, description string) error
	//UploadDirs git upload compress file
	//dir compress dir
	//filename  File with name only
	//description description
	UploadDirs(dir []string, filename string, description string) error
}

type BaseUpload struct {
	CompressType string
}

func (base *BaseUpload) GetCompress() (compress util.Compress, err error) {
	var ok bool
	if compress, ok = util.CompressSuports[base.CompressType]; !ok {
		err = errors.New("Compress type not Suport")
	}
	return
}
