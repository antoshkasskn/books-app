package http

import (
	"book-app/internal/entity"
	"book-app/pkg/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(port int, bookLogic entity.BookLogic) {
	var g = gin.New()

	g.Use(
		gin.Recovery(),
		middlewares.Logger(),
	)
	registerHealthCheck(g.Group("/health"))
	registerApiHandlers(g.Group("/api"), bookLogic)

	if err := g.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("can't start http server by error: %v", err)
	}
}
