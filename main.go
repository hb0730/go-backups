package main

import "github.com/hb0730/go-backups/cron"

func main() {
	err := cron.StartCron()
	panic(err)
}
