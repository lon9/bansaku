package main

import (
	"github.com/Rompei/zepher/controllers"
	"github.com/ipfans/echo-pongo2"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
)

// Routes is type of handlers
type Routes map[string]http.Handler

func (h Routes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := h[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// NewRoutes return routes
func NewRoutes() Routes {
	routes := make(Routes)

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
	routes["bansaku.localhost:8080"] = bansaku

	blog := echo.New()
	blog.Use(mw.Logger())
	blog.Use(mw.Recover())
	blog.SetDebug(true)
	blog.Static("/js", "blog/public/js")
	blog.Static("/css", "blog/public/css")
	blog.Static("/images", "blog/public/images")
	blog.Static("/font", "blog/public/font")
	blog.Static("/categories", "blog/public/categories")
	blog.Static("/page", "blog/public/page")
	blog.Static("/tags", "blog/public/tags")
	blog.Index("blog/public/index.html")
	routes["blog.localhost:8080"] = blog

	return routes
}
