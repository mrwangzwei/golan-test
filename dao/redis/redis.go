package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"self-test/config"
	"time"
)

var RedisClient *redis.Pool

func Init() {
	cfg := config.Conf
	RedisClient = &redis.Pool{
		MaxIdle:     cfg.RedisConfig.MaxIdle,   //最初的连接数量
		MaxActive:   cfg.RedisConfig.MaxActive, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second,         //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			addr := fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port)
			c, err := redis.Dial(cfg.RedisConfig.Type, addr)
			if err != nil {
				defer c.Close()
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}
			//if _, err := c.Do("AUTH", cfg.RedisConfig.Auth); err != nil {
			//	defer c.Close()
			//	fmt.Fprint(os.Stderr, err)
			//	os.Exit(1)
			//}
			return c, nil
		},
	}
}

func GetRedis() redis.Conn {
	return RedisClient.Get()
}
