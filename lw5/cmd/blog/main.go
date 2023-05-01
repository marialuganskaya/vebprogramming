package main

import (
	"log"
	"net/http"
)

const (
	port = ":3000"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/home", index)
	mux.HandleFunc("/post", post)

	log.Println("Start server " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
