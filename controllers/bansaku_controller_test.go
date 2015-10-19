package controllers

import (
	p "github.com/Rompei/zepher-bansaku/libs"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBansakuIndex(t *testing.T) {
	router := echo.New()
	router.Static("/js", "static/js")
	router.Static("/css", "static/css/bansaku")
	router.Static("/sound", "static/sound")
	router.Static("/font", "static/font")
	tmp := p.PrepareTemplates(p.Options{
		Directory:  "../templates/",
		Extensions: []string{".tpl"},
	})
	router.SetRenderer(tmp)
	server := NewBansakuServer()
	go server.Start()
	router.Get("/", BansakuIndex)
	router.WebSocket("/ws", server.BansakuSocketHandler())

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("URL is not found.")
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Status code is wrong.")
	}
}
