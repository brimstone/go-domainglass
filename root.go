package main

import "net/http"

func init() {
	http.Handle("/", Root())
}

// Root returns http filesystem
func Root() http.Handler {
	return http.StripPrefix("/", http.FileServer(http.Dir("root")))
}
