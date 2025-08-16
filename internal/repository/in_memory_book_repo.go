package repository

import (
	"book-app/internal/entity"
	"context"
	"sync"
)

type inMemoryBookRepository struct {
	storage *sync.Map
}

func NewBookRepo() entity.BookRepo {
	return &inMemoryBookRepository{
		storage: new(sync.Map),
	}
}

func (r *inMemoryBookRepository) Create(ctx context.Context, book *entity.Book) error {
	r.storage.Store(book.Id, book)
	return nil
}

func (r *inMemoryBookRepository) GetById(ctx context.Context, s string) (*entity.Book, error) {
	book, ok := r.storage.Load(s)
	if !ok {
		return nil, entity.ErrNotFound{}
	}
	return book.(*entity.Book), nil
}

func (r *inMemoryBookRepository) GetList(ctx context.Context) (books []*entity.Book, err error) {
	r.storage.Range(func(k, v interface{}) bool {
		books = append(books, v.(*entity.Book))
		return true
	})
	return
}

func (r *inMemoryBookRepository) DeleteById(ctx context.Context, s string) error {
	r.storage.Delete(s)
	return nil
}
