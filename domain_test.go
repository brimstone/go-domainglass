package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/brimstone/go-domainglass"
)

func Test_Domain(*testing.T) {
	m.InitEngine()
	ts := httptest.NewServer(m.Mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/v1")
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
