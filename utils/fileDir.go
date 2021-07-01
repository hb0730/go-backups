package util

import (
	"os"
	"path"
)

//GetTempDir temporary  Dir
func GetTempDir(pattern string) (string, error) {
	dir := os.TempDir()
	return os.MkdirTemp(dir, pattern)
}

//RemoveAll remove dir
func RemoveAll(dist string) error {
	return os.RemoveAll(dist)
}

//ExistExt file extension  isExist
func ExistExt(filename string) bool {
	ext := path.Ext(filename)
	if ext == "" {
		return false
	} else {
		return true
	}
}
