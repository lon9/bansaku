package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func BansakuIndex(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
