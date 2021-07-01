package uploads

type Upload interface {
	//UploadDir git upload compress file
	//dir compress dir
	//filename file name
	//description description
	UploadDir(dir, filename string, description string) error
	//UploadDirs git upload compress file
	//dir compress dir
	//filename file name
	//description description
	UploadDirs(dir []string, filename string, description string) error
}
