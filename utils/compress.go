package util

import (
	"compress/flate"
	"fmt"
	"github.com/mholt/archiver"
)

const (
	ZIP    = "zip"
	TAR    = "tar"
	TAR_GZ = "tar.gz"
	GZ     = "gz"
)

type Compress interface {
	CompressDir(dst string, src ...string) error
	CompressDirs(dest string, src []string) error
	Close() error
}

type ZipUtil struct {
	Zip *archiver.Zip
}

func NewZipUtil(compressionLevel int) *ZipUtil {
	zip := &archiver.Zip{
		CompressionLevel:     compressionLevel,
		OverwriteExisting:    true,
		MkdirAll:             true,
		SelectiveCompression: true,
		ContinueOnError:      true,
	}
	z := new(ZipUtil)
	z.Zip = zip
	return z
}
func NewZipUtilDefault() *ZipUtil {
	return NewZipUtil(flate.DefaultCompression)
}

func (zip *ZipUtil) CompressDir(dest string, src ...string) error {
	dest = withExtension(dest, ZIP)
	return zip.Zip.Archive(src, dest)
}
func (zip ZipUtil) CompressDirs(dest string, src []string) error {
	dest = withExtension(dest, ZIP)
	return zip.Zip.Archive(src, dest)
}

func (zip *ZipUtil) Close() error {
	return zip.Zip.Close()
}

type TarUtil struct {
	Tar *archiver.Tar
}

func NewTarUtil() *TarUtil {
	t := new(TarUtil)
	tar := &archiver.Tar{
		OverwriteExisting: true,
		MkdirAll:          true,
		ContinueOnError:   true,
	}
	t.Tar = tar
	return t
}

func (tar *TarUtil) CompressDir(dest string, src ...string) error {
	dest = withExtension(dest, TAR)
	return tar.Tar.Archive(src, dest)
}
func (tar *TarUtil) CompressDirs(dest string, src []string) error {
	dest = withExtension(dest, TAR)
	return tar.Tar.Archive(src, dest)
}
func (tar *TarUtil) Close() error {
	return tar.Tar.Close()
}

type TarGzUtil struct {
	GZ *archiver.TarGz
}

func NewTarGzUtil(compressLevel int) *TarGzUtil {
	gz := &archiver.TarGz{
		CompressionLevel: compressLevel,
		Tar: &archiver.Tar{
			OverwriteExisting: true,
			MkdirAll:          true,
			ContinueOnError:   true,
		},
	}
	return &TarGzUtil{GZ: gz}
}

func NewTarGzUtilDefault() *TarGzUtil {
	return NewTarGzUtil(flate.DefaultCompression)
}

func (gz *TarGzUtil) CompressDir(dest string, src ...string) error {
	dest = withExtension(dest, TAR_GZ)
	return gz.GZ.Archive(src, dest)
}
func (gz *TarGzUtil) CompressDirs(dest string, src []string) error {
	dest = withExtension(dest, TAR_GZ)
	return gz.GZ.Archive(src, dest)
}
func (gz *TarGzUtil) Close() error {
	return gz.GZ.Close()
}

func withExtension(dest, extension string) string {
	if !ExistExt(dest) {
		return fmt.Sprintf("%s.%s", dest, extension)
	}
	return dest
}

var CompressSuports = make(map[string]Compress)

func init() {
	CompressSuports[TAR] = NewTarUtil()
	CompressSuports[ZIP] = NewZipUtilDefault()
	CompressSuports[TAR_GZ] = NewTarGzUtilDefault()
}
