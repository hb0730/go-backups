package main

import (
	"github.com/hb0730/go-backups/config"
	"github.com/robfig/cron/v3"
)

func StartCron() error {
	c := cron.New(cron.WithSeconds())
	yaml := config.ReadYaml()
	r := yaml.StringMap("cron")
	for k, v := range r {
		job := nameJob{name: k}
		_, err := c.AddJob(v, job)
		if err != nil {
			panic(err)
		}
	}
	c.Run()
	defer c.Stop()
	return nil
}

type nameJob struct {
	name string
}

func (j nameJob) Run() {
	t := config.ReadYaml().String(j.name + ".type")
	switch t {
	case "git":
		var git Git
		err := config.ReadYaml().Unmarshal(j.name+".git", &git)
		if err != nil {
			panic(err)
		}
		filename := config.ReadYaml().String(j.name + ".uploads.filename")
		dirs := config.ReadYaml().Strings(j.name + ".uploads.dirs")
		upload(git, filename, dirs)
		break
	default:
		break
	}
}

type Git struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Dirpath  string `json:"dirpath"`
}

func upload(g Git, filename string, dir []string) {

}
