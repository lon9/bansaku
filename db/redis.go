package db

import (
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn
var isConnected bool

// GetRedis returns connection of redis.
func GetRedis() redis.Conn {
	if !isConnected {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
		isConnected = true
		return c
	}
	return c
}
