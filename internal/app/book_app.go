package app

import (
	"book-app/internal/http"
	"book-app/internal/logic"
	"book-app/internal/repository"
)

const httpPort = 8080

func Run() {
	bookRepo := repository.NewBookRepo()
	bookLogic := logic.NewBookLogic(bookRepo)
	http.Run(httpPort, bookLogic)
}
