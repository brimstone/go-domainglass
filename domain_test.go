package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	dg "github.com/brimstone/go-domainglass"
)

func Test_Domain(*testing.T) {
	ts := httptest.NewServer(dg.Mux)
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
