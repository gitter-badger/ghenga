package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlePerson(res http.ResponseWriter, req *http.Request) {
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/search/{searchTerm}", Search)
	// r.HandleFunc("/load/{dataId}", Load)
	r.HandleFunc("/api/person", HandlePerson)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
