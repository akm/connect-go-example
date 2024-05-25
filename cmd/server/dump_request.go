package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func DumpRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Dump the request.
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(string(dump))

		h.ServeHTTP(w, r)
	})
}
