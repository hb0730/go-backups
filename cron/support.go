package cron

var Supports map[string]Support

type Support interface {
	//Upload
	Upload(name, supportType string) error
}
