package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func init() {
	http.HandleFunc("/", Hello)
}

// Hello Handle main route
func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", runtime.Version())
}
