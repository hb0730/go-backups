package util

import (
	"os"
	"path"
	"strings"
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

//GetFilenameNoWithExtension get file name no with extension
func GetFilenameNoWithExtension(filepath string) string {
	filename := path.Base(filepath)
	ext := path.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

//GetFilenameWithExtension get filename with extension
func GetFilenameWithExtension(filepath string) string {
	return path.Base(filepath)
}
