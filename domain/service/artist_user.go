package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type ArtistUser interface {
	FindByID(string) (*entity.ArtistUser, error)
	Create(*entity.ArtistUser) (*entity.ArtistUser, error)
	Update(*entity.ArtistUser) (*entity.ArtistUser, error)
	Delete(*entity.ArtistUser) error
}

type artistUser struct {
	Repository repository.ArtistUser
}

func NewArtistUser(repo repository.ArtistUser) ArtistUser {
	return &artistUser{
		repo,
	}
}

func (s *artistUser) FindByID(id string) (*entity.ArtistUser, error) {
	return s.Repository.FindByID(id)
}

func (s *artistUser) Create(du *entity.ArtistUser) (*entity.ArtistUser, error) {
	return s.Repository.Store(du)
}

func (s *artistUser) Update(du *entity.ArtistUser) (*entity.ArtistUser, error) {
	return s.Repository.Update(du)
}

func (s *artistUser) Delete(du *entity.ArtistUser) error {
	return s.Repository.Delete(du)
}
