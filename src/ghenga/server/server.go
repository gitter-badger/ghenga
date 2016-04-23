package server

import (
	"encoding/json"
	"ghenga/db"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ListPeople handles listing person records.
func ListPeople(env *Env, res http.ResponseWriter, req *http.Request) error {
	people := []db.Person{}
	err := env.DbMap.Select(&people, "select * from people")
	if err != nil {
		return err
	}
	log.Printf("loaded %v person records", len(people))

	buf, err := json.Marshal(people)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)

	_, err = res.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

// ListenAndServe starts a new ghenga API server with the given environment.
func ListenAndServe(env *Env) (err error) {
	r := mux.NewRouter()

	// API routes
	r.Handle("/api/person", Handler{HandleFunc: ListPeople, Env: env}).Methods("GET")

	// server static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(env.Public)))

	// activate logging to stdout
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))

	return http.ListenAndServe(env.ListenAddr, nil)
}
