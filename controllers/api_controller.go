package controllers

import (
	"github.com/Rompei/zepher-bansaku/db"
	"github.com/Rompei/zepher-bansaku/models"
	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
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
