package util

import (
	"VideoStation/service"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	// 每日更新，cron is a job scheduler on Unix-like operating systems
	_, err := Cron.AddFunc("0 0 1 * * ?", service.ClicksStoreInDB)
	if err != nil {
		Logger().Info(err)
	}
	Cron.Start()
	Logger().Info("created cron job")
}
