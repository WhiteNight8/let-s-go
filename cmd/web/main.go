package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", "4000", "http network address")

	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/home/view", homeView)
	mux.HandleFunc("/home/create", homeCreate)

	log.Print("starting server on 4000")
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
