package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//GetTempDirFile tempdir/filename
func GetTempDirFile(pattern, filename string) (string, error) {
	dir, err := GetTempDir(pattern)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), filename), nil
}

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
func GetFilenameNoWithExtension(filename string) string {
	newFilename := filepath.Base(filename)
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(newFilename, ext)
}

//GetFilenameWithExtension get filename with extension
func GetFilenameWithExtension(filename string) string {
	return filepath.Base(filename)
}

//UploadFilenameWithDir upload dir and filename  y-m-d/filename
func UploadFilenameWithDir(filename string) string {
	return time.Now().Format("2006-01-02") + "/" + GetFilenameWithExtension(filename)
}
