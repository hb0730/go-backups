package uploads

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
