package repository

import (
	"github.com/tockn/emukone/domain/entity"
)

type ArtistUser interface {
	FindByID(string) (*entity.ArtistUser, error)
	Store(*entity.ArtistUser) (*entity.ArtistUser, error)
	Update(*entity.ArtistUser) (*entity.ArtistUser, error)
	Delete(*entity.ArtistUser) error
}
