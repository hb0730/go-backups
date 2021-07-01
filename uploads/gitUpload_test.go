package uploads

import (
	"testing"
)

const (
	url      = ""
	username = ""
	email    = ""
	token    = ""
)

func TestGitUpload_Clone(t *testing.T) {
	g := NewGitUpload(
		url,
		username,
		email,
		token,
		`E:\goWork\blog-backups`,
	)
	defer g.Close()
	err := g.Clone()
	if err != nil {
		panic(err)
	}
	err = g.UploadDir(`E:\goWork\zinx-demo`, `zinx-demo.zip`, "test")
	if err != nil {
		panic(err)
	}
}

func TestGitUpload_Close(t *testing.T) {
	g := NewGitUpload(
		url,
		username,
		email,
		token,
		`E:\goWork\blog-backups`,
	)
	err := g.Close()
	if err != nil {
		panic(err)
	}
}
