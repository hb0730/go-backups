package cron

import (
	"github.com/hb0730/go-backups/config"
	"github.com/hb0730/go-backups/uploads"
	util "github.com/hb0730/go-backups/utils"
)

type Git struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func init() {
	Supports["git"] = &Git{}
}

func (g *Git) Upload(name, supportKey string) error {
	//name: blog supportkey: git
	err := config.ReadYaml().Unmarshal(name+"."+supportKey, g)
	if err != nil {
		return err
	}

	return g.UploadFromGit(name)
}

func (g *Git) UploadFromGit(name string) error {
	filename := config.ReadYaml().String(name + ".uploads.filename")
	dirs := config.ReadYaml().Strings(name + ".uploads.dirs")

	dirPath, err := util.GetTempDir(filename)
	if err != nil {
		return err
	}
	git := uploads.NewGitUpload(g.URL, g.Username, g.Email, g.Email, dirPath)
	defer git.Close()
	return git.UploadDirs(dirs, filename, "auto backups")
}
