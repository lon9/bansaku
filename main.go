package main

import (
	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	routes := NewRoutes()
	gracehttp.Serve(routes.Server(":8080"))
}
