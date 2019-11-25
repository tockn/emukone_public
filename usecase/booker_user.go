package usecase

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type BookerUser interface {
	Show(string) (*entity.BookerUser, error)
	New(*entity.BookerUser) (string, error)
	Edit(*entity.BookerUser) (*entity.BookerUser, error)
	Delete(*entity.BookerUser) error
}

type bookerUser struct {
	bookerUserService service.BookerUser
	hashIDService     service.HashID
}

func NewBookerUser(u service.BookerUser, h service.HashID) BookerUser {
	return &bookerUser{
		u,
		h,
	}
}

func (u *bookerUser) Show(id string) (*entity.BookerUser, error) {
	return u.bookerUserService.FindByID(id)
}

func (u *bookerUser) New(du *entity.BookerUser) (string, error) {
	nu, err := u.bookerUserService.Create(du)
	if err != nil {
		return "", err
	}
	return nu.ID, nil
}

func (u *bookerUser) Edit(du *entity.BookerUser) (*entity.BookerUser, error) {
	return u.bookerUserService.Update(du)
}

func (u *bookerUser) Delete(du *entity.BookerUser) error {
	return u.bookerUserService.Delete(du)
}
