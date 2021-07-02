package cron

import (
	"github.com/hb0730/go-backups/config"
	"testing"
)

func TestGit_Upload(t *testing.T) {
	config.LoadKoanf("")
	g := Git{}
	err := g.Upload("blog", "git")
	if err != nil {
		panic(err)
	}
}
