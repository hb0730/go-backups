package util

import (
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
