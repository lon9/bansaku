package main

import (
	"github.com/tylerb/graceful"
	"time"
)

func main() {
	routes := NewRoutes()
	graceful.ListenAndServe(routes.Server(":60000"), 5*time.Second)
}
