package controllers

import (
	"github.com/Rompei/zepher-bansaku/db"
	"github.com/Rompei/zepher-bansaku/models"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"net"
	"net/http"
)

// APIBansakuGetHandler returns count of Bansaku
func APIBansakuGetHandler(c *echo.Context) error {
	con := db.GetRedis()
	count, err := redis.Int64(con.Do("get", "count"))
	if err != nil {
		count = 0
	}
	bansaku := models.Bansaku{
		Count: count,
	}
	return c.JSON(http.StatusOK, &bansaku)
}

func checkRateLimit(con redis.Conn, c *echo.Context) bool {
	ip, _, err := net.SplitHostPort(c.Request().RemoteAddr)
	if err != nil {
		panic(err)
	}
	current, err := redis.Int(con.Do("LLEN", ip))
	if err == nil && current > 10 {
		return false
	}
	_, err = redis.Bool(con.Do("EXISTS", "ip"))
	if err != nil {
		con.Send("MULTI")
		con.Send("RPUSH", "ip", ip)
		con.Send("EXPIRE", "ip", 1)
		_, err := con.Do("EXEC")
		if err != nil {
			panic(err)
		}
	} else {
		con.Do("RPUSHX", "ip", ip)
	}
	return true

}
