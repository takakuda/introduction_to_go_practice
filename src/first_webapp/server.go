package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	server := http.Server{
		Addr: "127:0:0:1:8080",
		Handler: nil
	}
	http.ListenAndServeTLS("cert.pem", "key.pem")
}
