package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func init() {
	http.HandleFunc("/api/v1/", Domain)
}

// Hello Handle main route
func Domain(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", runtime.Version())
}
