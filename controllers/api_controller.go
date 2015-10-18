package controllers

import (
	"github.com/Rompei/zepher-bansaku/db"
	"github.com/Rompei/zepher-bansaku/models"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"net"
	"net/http"
)

const (
	// ReachedRateLimit is message, when reach api's rate limit.
	ReachedRateLimit = 900
)

// APIBansakuGetHandler returns count of Bansaku
func APIBansakuGetHandler(c *echo.Context) error {
	con := db.GetRedis()
	defer con.Close()
	if !checkRateLimit(con, c) {
		err := models.Error{
			Code:    ReachedRateLimit,
			Message: "Reached rate limit.",
		}
		return c.JSON(http.StatusBadRequest, err)
	}
	count, err := redis.Int64(con.Do("get", "count"))
	if err != nil {
		count = 0
	}
	bansaku := models.Bansaku{
		Count: count,
	}
	return c.JSON(http.StatusOK, &bansaku)
}

// Checking rate limit 10 request / 1 sec.
func checkRateLimit(con redis.Conn, c *echo.Context) bool {
	ip, _, err := net.SplitHostPort(c.Request().RemoteAddr)
	if err != nil {
		panic(err)
	}

	//If list of ip address's length is 10 retun false.
	current, err := redis.Int(con.Do("LLEN", ip))
	if err == nil && current > 10 {
		return false
	}
	exists, err := redis.Bool(con.Do("EXISTS", ip))
	if err != nil {
		panic(err)
	}
	if !exists {
		con.Send("MULTI")
		con.Send("RPUSH", ip, ip)
		con.Send("EXPIRE", ip, 10)
		_, err := con.Do("EXEC")
		if err != nil {
			panic(err)
		}
	} else {
		_, err := con.Do("RPUSHX", ip, ip)
		if err != nil {
			panic(err)
		}
	}
	return true
}
