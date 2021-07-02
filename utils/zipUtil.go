package util

import (
	"compress/flate"
	"github.com/mholt/archiver"
)

type ZipUtils struct {
	Zip archiver.Zip
}

func NewZipUtils() *ZipUtils {
	zip := archiver.Zip{
		CompressionLevel:     flate.DefaultCompression,
		OverwriteExisting:    true,
		MkdirAll:             true,
		SelectiveCompression: true,
		ContinueOnError:      true,
	}
	z := new(ZipUtils)
	z.Zip = zip
	return z
}
func (z ZipUtils) Close() error {
	return z.Zip.Close()
}

//CompressDir compress dir
//srcDir Target directory
//dest File with path
func (z *ZipUtils) CompressDir(srcDir string, dest string) error {
	if !ExistExt(dest) {
		dest += ".zip"
	}
	return z.Zip.Archive([]string{srcDir}, dest)
}

//CompressDirs compress dirs
//srcDirs Target directory
//dest File with path
func (z ZipUtils) CompressDirs(srcDirs []string, dest string) error {
	if !ExistExt(dest) {
		dest += ".zip"
	}
	return z.Zip.Archive(srcDirs, dest)
}
