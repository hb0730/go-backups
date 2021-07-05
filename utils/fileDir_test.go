package util

import (
	"fmt"
	"os"
	"testing"
)

func TestGetTempDir(t *testing.T) {
	f, err := GetTempDir("")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(f)
}

func TestGetFilenameWithExtension(t *testing.T) {
	filename := GetFilenameWithExtension(`C:\Users\12780\AppData\Local\Temp\2964299587\zinx-demo.zip`)
	fmt.Println(filename)
}

func TestGetFilenameNoWithExtension(t *testing.T) {
	filename := GetFilenameNoWithExtension(`C:\Users\12780\AppData\Local\Temp\2964299587\zinx-demo.zip`)
	fmt.Println(filename)
}
