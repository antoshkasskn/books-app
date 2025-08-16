package logic

import (
	"book-app/internal/entity"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBookLogic_Create(t *testing.T) {
	mockRepo := new(MockBookRepo)
	logic := NewBookLogic(mockRepo)

	book := &entity.Book{Title: "Test Book", Author: "Author Name"}

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.Book")).Return(nil)

	err := logic.Create(context.Background(), book)

	assert.NoError(t, err)
	assert.NotEmpty(t, book.Id) // Проверяем, что ID был установлен
	mockRepo.AssertExpectations(t)
}

func TestBookLogic_GetById(t *testing.T) {
	mockRepo := new(MockBookRepo)
	logic := NewBookLogic(mockRepo)

	book := &entity.Book{Id: "1", Title: "Test Book", Author: "Author Name"}
	mockRepo.On("GetById", mock.Anything, "1").Return(book, nil)

	result, err := logic.GetById(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, book, result)
	mockRepo.AssertExpectations(t)
}

func TestBookLogic_GetList(t *testing.T) {
	mockRepo := new(MockBookRepo)
	logic := NewBookLogic(mockRepo)

	books := []*entity.Book{
		{Id: "1", Title: "Test Book 1", Author: "Author 1"},
		{Id: "2", Title: "Test Book 2", Author: "Author 2"},
	}
	mockRepo.On("GetList", mock.Anything).Return(books, nil)

	result, total, err := logic.GetList(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, books, result)
	assert.Equal(t, total, len(books))
	mockRepo.AssertExpectations(t)
}

func TestBookLogic_DeleteById(t *testing.T) {
	mockRepo := new(MockBookRepo)
	logic := NewBookLogic(mockRepo)

	mockRepo.On("DeleteById", mock.Anything, "1").Return(nil)

	err := logic.DeleteById(context.Background(), "1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBookLogic_DeleteById_Error(t *testing.T) {
	mockRepo := new(MockBookRepo)
	logic := NewBookLogic(mockRepo)

	mockRepo.On("DeleteById", mock.Anything, "1").Return(errors.New("not found"))

	err := logic.DeleteById(context.Background(), "1")

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	mockRepo.AssertExpectations(t)
}
