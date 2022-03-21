package cache

import (
	"VideoStation/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

package main

import (
"context"
"fmt"
"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	ctx    context.Context
)

func InitConn() {

	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisSetting.Host,
		Password: conf.RedisSetting.Password,
		DB:       0,
		IdleTimeout: conf.RedisSetting.IdleTimeout,
	})

	ctx = context.Background()

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis 连接失败", err.Error())
		return
	}
	fmt.Println("redis 连接成功", ping)
}

func SetValue(key string, value interface{}) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
}

func GetValue(key string) {
	val, err := client.Get(ctx, key).Result()
	if err != redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err)
		//panic(err)
	} else {
		fmt.Println(key, val)
	}

}

