package main

import "net/http"

func init() {
	http.Handle("/assets/", Assets())
}

// Assets returns http filesystem
func Assets() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
}
