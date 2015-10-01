package main

import (
	"github.com/Rompei/zepher/controllers"
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
	bansaku.Use(mw.Logger())
	bansaku.Use(mw.Recover())
	// Debug
	bansaku.SetDebug(true)
	bansaku.Get("/", controllers.BansakuIndex)
	routes["bansaku.localhost:1323"] = bansaku

	return routes
}
