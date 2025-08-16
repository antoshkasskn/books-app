package http

import (
	"book-app/internal/entity"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	version1 = "/v1"
	idParam  = "id"
)

type apiHandler struct {
	bookLogic entity.BookLogic
}

var bookValidator = validator.New()

func registerApiHandlers(g *gin.RouterGroup, bookLogic entity.BookLogic) {
	h := &apiHandler{
		bookLogic: bookLogic,
	}
	v1 := g.Group(version1)

	v1.POST("/books", h.createBook)
	v1.GET("/books", h.getBooks)
	v1.GET("/books/:"+idParam, h.getBookById)
	v1.DELETE("/books/:"+idParam, h.deleteById)
}

func (h *apiHandler) createBook(ctx *gin.Context) {
	var book *entity.Book

	err := ctx.Bind(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bookValidator.Struct(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.bookLogic.Create(ctx, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func (h *apiHandler) getBooks(ctx *gin.Context) {
	books, total, err := h.bookLogic.GetList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  books,
		"total": total,
	})
}

func (h *apiHandler) getBookById(ctx *gin.Context) {
	id := ctx.Param(idParam)
	books, err := h.bookLogic.GetById(ctx, id)
	if err != nil {
		if errors.As(err, &entity.ErrNotFound{}) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Book not found by id %s", id)})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (h *apiHandler) deleteById(ctx *gin.Context) {
	id := ctx.Param(idParam)
	err := h.bookLogic.DeleteById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
