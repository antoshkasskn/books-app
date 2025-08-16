package http

import (
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle request %s %s from %s %s", r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent())
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("HEALTHY"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
