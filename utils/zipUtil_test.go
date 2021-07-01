package util

import (
	"testing"
)

func TestNewZipUtilsFromPath(t *testing.T) {
	z := NewZipUtils()
	err := z.CompressDir(`E:\goWork\zinx-demo`, `E:\goWork\zinx-demo`)
	if err != nil {
		panic(err)
	}
}
