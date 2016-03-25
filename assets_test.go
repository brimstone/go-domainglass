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

func Test_Assets(*testing.T) {
	ts := httptest.NewServer(m.Assets())
	defer ts.Close()

	fmt.Println(ts.URL + "/assets/style.css")
	res, err := http.Get(ts.URL + "/assets/style.css")
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
