package usecase

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type ArtistUser interface {
	Show(string) (*entity.ArtistUser, error)
	New(*entity.ArtistUser) (string, error)
	Edit(*entity.ArtistUser) (*entity.ArtistUser, error)
	Delete(*entity.ArtistUser) error
}

type artistUser struct {
	artistUserService service.ArtistUser
	hashIDService     service.HashID
}

func NewArtistUser(u service.ArtistUser, h service.HashID) ArtistUser {
	return &artistUser{
		u,
		h,
	}
}

func (u *artistUser) Show(id string) (*entity.ArtistUser, error) {
	return u.artistUserService.FindByID(id)
}

func (u *artistUser) New(du *entity.ArtistUser) (string, error) {
	nu, err := u.artistUserService.Create(du)
	if err != nil {
		return "", err
	}
	return nu.ID, nil
}

func (u *artistUser) Edit(du *entity.ArtistUser) (*entity.ArtistUser, error) {
	return u.artistUserService.Update(du)
}

func (u *artistUser) Delete(du *entity.ArtistUser) error {
	return u.artistUserService.Delete(du)
}
