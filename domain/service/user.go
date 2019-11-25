package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type User interface {
	FindMetaByID(string) (*entity.UserMeta, error)
	FindIDByExternalService(*entity.ExternalService) (string, bool, error)
	NewExternalService(*entity.ExternalService, string) error
	FindMetaByEventID(string) ([]*entity.UserMeta, error)
	FindTag() ([]*entity.UserTag, error)
}

type user struct {
	Repository repository.User
}

func NewUser(r repository.User) User {
	return &user{
		r,
	}
}

func (s *user) FindMetaByID(id string) (*entity.UserMeta, error) {
	return s.Repository.FindMetaByID(id)
}

func (s *user) FindIDByExternalService(service *entity.ExternalService) (string, bool, error) {
	return s.Repository.FindIDByExternalService(service)
}

func (s *user) NewExternalService(es *entity.ExternalService, userID string) error {
	return s.Repository.StoreExternalService(es, userID)
}

func (s *user) FindMetaByEventID(id string) ([]*entity.UserMeta, error) {
	return s.Repository.FindMetaByEventID(id)
}

func (s *user) FindTag() ([]*entity.UserTag, error) {
	return s.Repository.FindTag()
}
