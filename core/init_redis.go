package core

import (
	"boke-server/global"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var DB *redis.Client

func InitRedis() *redis.Client {
	r := global.Config.Redis
	fmt.Println(r)
	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatalf("redis连接失败", err)
	}
	fmt.Println("redis链接成功")
	return redisDB
}
