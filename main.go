package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hello from xia"))
}

func homeView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific view"))
}

func homeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a specific view"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/home/view", homeView)
	mux.HandleFunc("/home/create", homeCreate)

	log.Print("starting server on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
