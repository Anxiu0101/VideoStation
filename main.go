package main

import (
	"VideoStation/models"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/setting"
	"VideoStation/routers"
	"fmt"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
