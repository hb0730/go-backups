package util

import (
	"fmt"
	"testing"
)

var (
	dest = `E:\goWork\test`
	src  = []string{`E:\goWork\zinx-demo`, `E:\goWork\go-otp`}
)

func TestZipUtil_CompressDir(t *testing.T) {
	zip := NewZipUtilDefault()
	defer zip.Close()
	err := zip.CompressDir(dest, src...)
	if err != nil {
		panic(err)
	}
}

func withExtensionTest(dest string, extension string) string {
	if !ExistExt(dest) {
		return fmt.Sprintf("%s.%s", dest, extension)
	}
	return dest
}

func Test_withExtension(t *testing.T) {
	dest = withExtensionTest(dest, "zip")
	fmt.Println(dest)
}
