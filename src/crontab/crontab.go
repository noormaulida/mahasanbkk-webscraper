package crontab

import (
    "mahasanbkk-webscraper/src/service"

    cron "github.com/robfig/cron/v3"
)

func ApplyScheduler() {
	scheduler := cron.New()
	defer scheduler.Stop()

	scheduler.AddFunc("*/1 * * * *", service.WebScraper)
	go scheduler.Start()
}

