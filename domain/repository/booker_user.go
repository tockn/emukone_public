package repository

import (
	"github.com/tockn/emukone/domain/entity"
)

type BookerUser interface {
	FindByID(string) (*entity.BookerUser, error)
	Store(*entity.BookerUser) (*entity.BookerUser, error)
	Update(*entity.BookerUser) (*entity.BookerUser, error)
	Delete(*entity.BookerUser) error
}
