package entity

import (
	"context"
	"time"
)

type BookLogic interface {
	Create(context.Context, *Book) error
	GetById(context.Context, string) (*Book, error)
	GetList(context.Context) ([]*Book, int, error)
	DeleteById(context.Context, string) error
}

type BookRepo interface {
	Create(context.Context, *Book) error
	GetById(context.Context, string) (*Book, error)
	GetList(context.Context) ([]*Book, int, error)
	DeleteById(context.Context, string) error
}

type Book struct {
	Id              string    `json:"id" validate:"omitempty,min=1,max=255" db:"id"`
	Title           string    `json:"title" validate:"required,min=1,max=255" db:"title"`
	Author          string    `json:"author" validate:"required,min=1,max=128" db:"author"`
	PublicationDate time.Time `json:"publicationDate" validate:"required" db:"publication_date"`
	Publisher       string    `json:"publisher" validate:"required,min=1,max=128" db:"publisher"`
	Edition         int       `json:"edition" validate:"required,min=1" db:"edition"`
	Location        string    `json:"location" validate:"required,min=1,max=128" db:"location"`
}
