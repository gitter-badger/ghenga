// Package server contains the API server implementation and base
// functionality for ghenga.
package server

import (
	"encoding/json"
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

type jsonError struct {
	Error string `json:"error"`
}

// ServeHTTP allows the handler to be used in place of http.Handler.
func (h Handler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	err := h.HandleFunc(h.Env, wr, req)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// return the error to the client as a nicely formatted json document.
			buf, err := json.Marshal(jsonError{Error: e.Error()})
			if err != nil {
				panic(err)
			}
			wr.Header().Set("Content-Type", "application/json")
			wr.WriteHeader(e.Status())
			wr.Write(buf)
		default:
			// return a generic internal server error message with status 500
			http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
