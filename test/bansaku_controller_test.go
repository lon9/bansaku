package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_BansakuIndex(t *testing.T) {
	router := NewRouter()
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
