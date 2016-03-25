package main_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/brimstone/go-domainglass"
)

func Test_Hello(*testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(m.Hello))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, "expected 200")
	}
}
