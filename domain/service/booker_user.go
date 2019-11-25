package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type BookerUser interface {
	FindByID(string) (*entity.BookerUser, error)
	Create(*entity.BookerUser) (*entity.BookerUser, error)
	Update(*entity.BookerUser) (*entity.BookerUser, error)
	Delete(*entity.BookerUser) error
}

type bookerUser struct {
	Repository repository.BookerUser
}

func NewBookerUser(repo repository.BookerUser) BookerUser {
	return &bookerUser{
		repo,
	}
}

func (s *bookerUser) FindByID(id string) (*entity.BookerUser, error) {
	return s.Repository.FindByID(id)
}

func (s *bookerUser) Create(du *entity.BookerUser) (*entity.BookerUser, error) {
	return s.Repository.Store(du)
}

func (s *bookerUser) Update(du *entity.BookerUser) (*entity.BookerUser, error) {
	return s.Repository.Update(du)
}

func (s *bookerUser) Delete(du *entity.BookerUser) error {
	return s.Repository.Delete(du)
}
