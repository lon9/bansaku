package main

import (
	"github.com/Rompei/zepher-bansaku/controllers"
	"github.com/ipfans/echo-pongo2"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// NewRoutes return routes
func NewRoutes() *echo.Echo {
	bansaku := echo.New()
	bansaku.Static("/js", "static/js")
	bansaku.Static("/sound", "static/sound")
	bansaku.Use(mw.Logger())
	bansaku.Use(mw.Recover())
	bansaku.Use(pongo2.Pongo2())
	// Debug
	bansaku.SetDebug(true)
	server := controllers.NewBansakuServer()
	go server.Start()
	bansaku.Get("/", controllers.BansakuIndex)
	bansaku.WebSocket("/ws", server.BansakuSocketHandler())

	return bansaku
}
