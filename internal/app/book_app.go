package app

import (
	"book-app/internal/http"
	"book-app/internal/logic"
	"book-app/internal/repository"
	"book-app/pkg/config"
	"book-app/pkg/logger"
	"github.com/jackc/pgx"
	"log"
)

const httpPort = 8080

func Run() {
	logger.InitLogger()
	pgPool, err := pgx.NewConnPool(config.Cfg.DbConfig.GetPgxConf())
	if err != nil {
		log.Fatal(err)
	}
	bookRepo := repository.NewBookRepo(pgPool)
	bookLogic := logic.NewBookLogic(bookRepo)
	http.Run(httpPort, bookLogic)
}
