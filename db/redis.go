package db

import (
	"github.com/garyburd/redigo/redis"
)

// Redis is wrapper of redis.Conn
type Redis struct {
	Con redis.Conn
}

var sharedInstance = newRedis()

func newRedis() *Redis {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return &Redis{
		Con: c,
	}
}

// GetInstance of Redis.
func GetInstance() *Redis {
	return sharedInstance
}
