package redis

import (
	//"os"
	//"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func Client() *redis.Pool {

	var pool *redis.Pool

	pool = &redis.Pool{
		MaxIdle:     120000,
		MaxActive:   0,
		IdleTimeout: 2 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.DialTimeout("tcp", "localhost:6379", 100*time.Millisecond, 100*time.Millisecond, 100*time.Millisecond)
			if err != nil {
				return nil, err
			}
			return conn, err
		},
	}

	return pool
}

func Set(pool *redis.Pool, key string, val string) string {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val, "EX", 120)
	if err != nil {
		return "error"
	}
	return "ok"
}

func Get(pool *redis.Pool, key string) string {
	conn := pool.Get()
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "error"
	} else {
		return val
	}
}
