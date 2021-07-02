package cron

var Supports = make(map[string]Support)

type Support interface {
	//Upload
	Upload(name, supportType string) error
}
