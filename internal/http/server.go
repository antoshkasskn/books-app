package http

import (
	"fmt"
	"log"
	"net/http"
)

func Run(port int) {
	var (
		mux    = http.NewServeMux()
		server = http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		}
	)

	mux.HandleFunc("GET /health", health)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("can't start http server by error: %v", err)
	}
}
