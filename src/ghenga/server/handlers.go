// Package server contains the API server implementation and base
// functionality for ghenga.
package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/modl"
)

// Env is an environment for a handler function.
type Env struct {
	DbMap      *modl.DbMap
	ListenAddr string
	Public     string
}

// Handler is an http.Handler with an explicit error return value, bundled together with an environment.
type Handler struct {
	Env        *Env
	HandleFunc func(*Env, http.ResponseWriter, *http.Request) error
}

// httpWriteJSON encodes the given structures as JSON and writes them to the
// ResponseWriter.
func httpWriteJSON(wr http.ResponseWriter, status int, data interface{}) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(status)

	enc := json.NewEncoder(wr)

	return enc.Encode(data)
}

// ServeHTTP allows the handler to be used in place of http.Handler.
func (h Handler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	err := h.HandleFunc(h.Env, wr, req)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// return the error to the client as a nicely formatted json document.

			type jsonError struct {
				Error string `json:"error"`
			}

			err = httpWriteJSON(wr, e.Status(), jsonError{Error: e.Error()})
			if err != nil {
				log.Printf("error writing error document to client: %v", err)
			}
		default:
			// return a generic internal server error message with status 500
			http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
