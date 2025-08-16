package logic

import (
	"book-app/internal/entity"
	"context"
	"github.com/google/uuid"
)

type bookLogic struct {
	bookRepo entity.BookRepo
}

func NewBookLogic(bookRepo entity.BookRepo) entity.BookLogic {
	return &bookLogic{
		bookRepo: bookRepo,
	}
}

func (logic *bookLogic) Create(ctx context.Context, book *entity.Book) error {
	book.Id = uuid.New().String()
	return logic.bookRepo.Create(ctx, book)
}

func (logic *bookLogic) GetById(ctx context.Context, id string) (*entity.Book, error) {
	return logic.bookRepo.GetById(ctx, id)
}

func (logic *bookLogic) GetList(ctx context.Context) ([]*entity.Book, int, error) {
	return logic.bookRepo.GetList(ctx)
}

func (logic *bookLogic) DeleteById(ctx context.Context, i string) error {
	return logic.bookRepo.DeleteById(ctx, i)
}
