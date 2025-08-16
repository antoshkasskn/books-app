package app

import "book-app/internal/http"

const httpPort = 8080

func Run() {
	http.Run(httpPort)
}
