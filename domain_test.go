package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	dg "github.com/brimstone/go-domainglass"
)

func Test_NotBeta(*testing.T) {
	ts := httptest.NewServer(dg.Mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/domain.glass")
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, " expected 200")
	}
}

func Test_Beta(*testing.T) {
	var err error
	ts := httptest.NewServer(dg.Mux)
	defer ts.Close()

	// setup beta
	res, err := http.Get(ts.URL + "/beta")
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, " expected 200")
	}
}

func Test_Domain(*testing.T) {
	ts := httptest.NewServer(dg.Mux)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/domain.glass", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", "beta=true")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, " expected 200")
	}
}

func Test_API(*testing.T) {
	var err error
	ts := httptest.NewServer(dg.Mux)
	defer ts.Close()

	// setup beta
	res, err := http.Get(ts.URL + "/api/v1/domain.glass")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, " expected 200")
	}
	fmt.Println(string(body))
}

func Test_APINew(*testing.T) {
	var err error
	ts := httptest.NewServer(dg.Mux)
	defer ts.Close()

	// setup beta
	res, err := http.Post(
		ts.URL+"/api/v1/new",
		"application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte("domain=domain.glass")),
	)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("Status code is ", res.Status, " expected 200")
	}
	fmt.Println(string(body))
}
