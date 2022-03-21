package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//func init() {
//	conf.Setup()
//	models.Setup()
//	logging.Setup()
//}

//func main() {
//	router := routers.InitRouter()
//
//	s := &http.Server{
//		Addr:           fmt.Sprintf(":%d", conf.ServerSetting.HttpPort),
//		Handler:        router,
//		ReadTimeout:    conf.ServerSetting.ReadTimeout,
//		WriteTimeout:   conf.ServerSetting.WriteTimeout,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	s.ListenAndServe()
//}

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}
