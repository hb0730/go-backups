package cron

import (
	"errors"
	"github.com/hb0730/go-backups/config"
	"github.com/mritd/logger"
	"github.com/robfig/cron/v3"
)

func StartCron() error {
	c := cron.New(cron.WithSeconds())
	yaml := config.ReadYaml()
	r := yaml.StringMap("cron")
	if len(r) == 0 {
		return errors.New("cron size is 0")
	}
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
	support := config.ReadYaml().String(j.name + ".type")
	logger.Info("[cron] ", "job run ", "support: ", support)
	if v, ok := Supports[support]; ok {
		//blog.git
		//blog.aliyun-oss
		err := v.Upload(j.name, support)
		if err != nil {
			logger.Errorf("upload support error: %s", err.Error())
		}
	}
}
