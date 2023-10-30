package handlers

import (
	"log"
	"net/http"
)

func LogHandler(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host + " " + r.Method + " " + r.URL.Path + " " + r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
