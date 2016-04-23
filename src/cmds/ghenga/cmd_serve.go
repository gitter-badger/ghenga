package main

import (
	"encoding/json"
	"ghenga"
	"ghenga/db"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/modl"
)

var DBM *modl.DbMap

// ListPeople handles listing person records.
func ListPeople(res http.ResponseWriter, req *http.Request) (code int, err error) {
	people := []db.Person{}
	err = DBM.Select(&people, "select * from people")
	if err != nil {
		return 0, err
	}
	log.Printf("loaded %v person records", len(people))

	buf, err := json.Marshal(people)
	if err != nil {
		return 0, err
	}

	_, err = res.Write(buf)
	if err != nil {
		return 0, err
	}

	return http.StatusOK, nil
}

type cmdServe struct {
	Port uint `short:"p" long:"port" default:"8080" description:"set the port for the HTTP server"`
}

func init() {
	_, err := parser.AddCommand("serve",
		"start server",
		"The server command starts the HTTP server",
		&cmdServe{})
	if err != nil {
		panic(err)
	}
}

func (opts *cmdServe) Execute(args []string) (err error) {
	var cleanup func() error
	DBM, cleanup, err = OpenDB()
	if err != nil {
		return err
	}
	defer CleanupErr(&err, cleanup)

	log.Printf("starting server at port %v", opts.Port)

	r := mux.NewRouter()
	r.HandleFunc("/api/person", ghenga.WrapError(ListPeople)).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
