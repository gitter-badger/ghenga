package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandlePerson(res http.ResponseWriter, req *http.Request) {

}

func HandleDefault(res http.ResponseWriter, req *http.Request) {
	log.Printf("default handler called")
}

func main() {
	dbmap, err := initDB("db/test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbmap.Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()


	r := mux.NewRouter()
	r.HandleFunc("/api/person", HandlePerson)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
