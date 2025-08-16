package logic

import (
	"book-app/internal/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

// MockBookRepo - мок для entity.BookRepo
type MockBookRepo struct {
	mock.Mock
}

func (m *MockBookRepo) Create(ctx context.Context, book *entity.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockBookRepo) GetById(ctx context.Context, id string) (*entity.Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *MockBookRepo) GetList(ctx context.Context) ([]*entity.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Book), args.Error(1)
}

func (m *MockBookRepo) DeleteById(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
