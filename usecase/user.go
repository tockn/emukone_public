package usecase

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type User interface {
	ShowMeta(string) (*entity.UserMeta, error)
	ShowIDByExternalService(*entity.ExternalService) (string, bool, error)
	NewExternalService(*entity.ExternalService, string) error
	ShowMetaByEventID(string) ([]*entity.UserMeta, error)
	ShowAllTags() ([]*entity.UserTag, error)
}

type user struct {
	userService   service.User
	hashIDService service.HashID
}

func NewUser(u service.User, h service.HashID) User {
	return &user{
		u,
		h,
	}
}

func (u *user) ShowMeta(id string) (*entity.UserMeta, error) {
	return u.userService.FindMetaByID(id)
}

func (u *user) ShowIDByExternalService(es *entity.ExternalService) (string, bool, error) {
	return u.userService.FindIDByExternalService(es)
}

func (u *user) NewExternalService(es *entity.ExternalService, id string) error {
	return u.userService.NewExternalService(es, id)

}

func (u *user) ShowMetaByEventID(id string) ([]*entity.UserMeta, error) {
	return u.userService.FindMetaByEventID(id)
}

func (u *user) ShowAllTags() ([]*entity.UserTag, error) {
	return u.userService.FindTag()
}
