package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
	log.Print(r.URL.Path)
}

func Error404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
	log.Print(r.URL.Path)
}

func main() {
	var srv http.Server

	srv.Addr = ":8000"
	http2.ConfigureServer(&srv, nil)
	
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/favicon.ico", Error404)
	log.Fatal(srv.ListenAndServeTLS("localhost.crt", "localhost.key"))
}
