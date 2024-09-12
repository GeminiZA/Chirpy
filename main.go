package main

import "net/http"

func root() {
}

func main() {
	mux := http.NewServeMux()

  mux.HandleFunc("/", ))
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
