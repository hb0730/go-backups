package main

import (
	"flag"
	"github.com/hb0730/go-backups/config"
	"github.com/hb0730/go-backups/cron"
	"github.com/mritd/logger"
)

var configfile string

func init() {
	flag.StringVar(&configfile, "c", "", "read config file")
}
func main() {
	flag.Parse()
	logger.Info("[main] ", "start success ", "properties config: ", configfile)
	config.LoadKoanf(configfile)
	err := cron.StartCron()
	panic(err)
}
