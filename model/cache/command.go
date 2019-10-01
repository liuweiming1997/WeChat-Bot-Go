package cache

import (
	"github.com/WeChat-Bot-Go/logger"
	"github.com/gomodule/redigo/redis"
)

func Get(key []byte) ([]byte, error) {
	conn := GlobalRedisPool.Get()
	defer conn.Close()

	res, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		logger.Error("Redis get error", err)
		return nil, err
	}
	return res, nil
}

func Set(key, value []byte) error {
	conn := GlobalRedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

func Exist(key []byte) bool {
	conn := GlobalRedisPool.Get()
	defer conn.Close()

	res, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		logger.Error("Redis exist error", err)
		return false
	}
	return res
}

func Delete(key []byte) error {
	conn := GlobalRedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func ClearAll() error {
	conn := GlobalRedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	return err
}
