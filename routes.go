package main

import (
	"github.com/Rompei/zepher-bansaku/controllers"
	p "github.com/Rompei/zepher-bansaku/libs"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// NewRoutes return routes
func NewRoutes() *echo.Echo {
	bansaku := echo.New()
	bansaku.Static("/js", "static/js")
	bansaku.Static("/css", "static/css")
	bansaku.Static("/sound", "static/sound")
	bansaku.Static("/font", "static/font")
	bansaku.Use(mw.Logger())
	bansaku.Use(mw.Recover())
	//bansaku.Use(pongo2.Pongo2())
	t := p.PrepareTemplates(p.Options{
		Directory:  "templates/",
		Extensions: []string{".tpl"},
	})
	bansaku.SetRenderer(t)

	// Debug
	bansaku.SetDebug(true)
	server := controllers.NewBansakuServer()
	go server.Start()
	bansaku.Get("/", controllers.BansakuIndex)
	bansaku.WebSocket("/ws", server.BansakuSocketHandler())

	// API
	api := bansaku.Group("/api")
	api.Get("/", controllers.APIReferenceHandler)
	api.Get("/count", controllers.APIBansakuGetHandler)

	return bansaku
}
