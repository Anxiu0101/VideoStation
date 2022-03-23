package cache

import (
	"VideoStation/conf"
	"VideoStation/pkg/util"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	Ctx         context.Context
)

func Setup() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        conf.RedisSetting.Host,
		Password:    conf.RedisSetting.Password,
		DB:          0,
		IdleTimeout: conf.RedisSetting.IdleTimeout,
	})

	Ctx = context.Background()

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		util.Logger().Info("redis fail to connect", err.Error())
		return
	}
}

func SetValue(key string, value interface{}) {
	err := RedisClient.Set(Ctx, key, value, 0).Err()
	if err != nil {
		util.Logger().Info(err)
		panic(err)
	}
}

func GetValue(key string) {
	val, err := RedisClient.Get(Ctx, key).Result()
	if err != redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		util.Logger().Info(err)
		panic(err)
	} else {
		fmt.Println(key, val)
	}

}
