package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	dg "github.com/brimstone/go-domainglass"
)

func TestMain(m *testing.M) {
	dg.InitEngine()
	dg.InitDatabase()
	os.Exit(m.Run())
}

func Test_Root(*testing.T) {
	ts := httptest.NewServer(dg.Mux)
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
	dg.GetBind()
}
