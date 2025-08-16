package entity

import (
	"context"
	"time"
)

type BookLogic interface {
	Create(context.Context, *Book) error
	GetById(context.Context, string) (*Book, error)
	GetList(context.Context) ([]*Book, error)
	DeleteById(context.Context, string) error
}

type BookRepo interface {
	Create(context.Context, *Book) error
	GetById(context.Context, string) (*Book, error)
	GetList(context.Context) ([]*Book, error)
	DeleteById(context.Context, string) error
}

type Book struct {
	Id              string    `json:"id" validate:"omitempty,min=1,max=255"`
	Title           string    `json:"title" validate:"required,min=1,max=255"`
	Author          string    `json:"author" validate:"required,min=1,max=128"`
	PublicationDate time.Time `json:"publicationDate" validate:"required"`
	Publisher       string    `json:"publisher" validate:"required,min=1,max=128"`
	Edition         int       `json:"edition" validate:"required,min=1"`
	Location        string    `json:"location" validate:"required,min=1,max=128"`
}
