package main

import (
	"github.com/tylerb/graceful"
	"time"
)

func main() {
	routes := NewRoutes()
	graceful.ListenAndServe(routes.Server(":8080"), 5*time.Second)
}
