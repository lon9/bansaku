package test

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_APIReferenceHandler(t *testing.T) {
	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/")
	if err != nil {
		t.Error("URL is not found.")
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Status code is wrong.")
	}

}

func Test_APIBansakuGetHandler(t *testing.T) {
	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/count")
	if err != nil {
		t.Error("URL is not foung.")
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Status code was wrong.")
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error("Can not parse body.")
	}
	js, err := simplejson.NewJson(b)
	if err != nil {
		t.Error("Can not get json.")
	}
	_, err = js.Get("count").Int64()
	if err != nil {
		t.Error("Can not get count.")
	}
}
