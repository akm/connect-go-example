package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type dumpRequest struct {
	Method string
	URL    string
	Proto  string
	Header http.Header
}

func DumpRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump := dumpRequest{
			Method: r.Method,
			URL:    r.URL.String(),
			Proto:  r.Proto,
			Header: r.Header,
		}
		b, err := json.Marshal(dump)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(string(b))

		h.ServeHTTP(w, r)
	})
}
