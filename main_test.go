package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/brimstone/go-domainglass"
)

func Test_Root(*testing.T) {
	m.InitEngine()
	m.InitDatabase()
	ts := httptest.NewServer(m.Mux)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, "expected 200")
	}
}

func Test_GetBind(*testing.T) {
	m.GetBind()
}
