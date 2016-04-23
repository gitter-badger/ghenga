// Package ghenga contains the API server implementation and base
// functionality.
package ghenga

import (
	"fmt"
	"net/http"
)

// Handler is an http.Handler with explicit return values.
type Handler func(http.ResponseWriter, *http.Request) (int, error)

// WrapError takes an ErrorHandler and returns an http.HandlerFunc.
func WrapError(h Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		code, err := h(res, req)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(res, "error: %v\n", err)
			return
		}

		res.WriteHeader(code)
	}
}
