package main

import (
	"net/http"
)

func main() {
	routes := NewRoutes()
	http.ListenAndServe(":1323", routes)
}
