// use redis for cache
package cache

import (
	"time"

	"github.com/WeChat-Bot-Go/conf"
	"github.com/WeChat-Bot-Go/logger"
	"github.com/gomodule/redigo/redis"
)

var GlobalRedisPool *redis.Pool

func init() {
	GlobalRedisPool = newPool(conf.GlobalConfig.REDIS_HOST)

	// go func() {
	//  ticker := time.NewTicker(time.Hour * 36)
	//  for t := range ticker.C {
	//    logrus.Info("time := ", t, "clear all")
	//    ClearAll()
	//  }
	// }()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,

		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			//TODO: find how to use it
			if time.Since(t) > 60*time.Second {
				logger.Info("redis do ping")
				_, err := c.Do("PING")
				return err
			}
			return nil
		},
	}
}
