package main

import (
	"net/http"
)

func main() {
	routes := NewRoutes()
	http.ListenAndServe(":8080", routes)
}
