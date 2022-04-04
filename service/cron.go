package service

import (
	"VideoStation/pkg/util"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}

	// 每日更新，cron is a job scheduler on Unix-like operating systems
	_, err := Cron.AddFunc("0 0 1 * * ?", ClicksStoreInDB)
	if err != nil {
		util.Logger().Info(err)
	}
	_, err = Cron.AddFunc("0 0 1 * * ?", HistoryStoreInDB)
	if err != nil {
		util.Logger().Info(err)
	}
	Cron.Start()
	util.Logger().Info("created cron job")
}
