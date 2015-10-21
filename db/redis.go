package db

import (
	"github.com/garyburd/redigo/redis"
)

// GetRedis returns connection of redis.
func GetRedis() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return c
}
