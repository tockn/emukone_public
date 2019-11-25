package repository

import (
	"github.com/tockn/emukone/domain/entity"
)

type User interface {
	FindMetaByID(string) (*entity.UserMeta, error)
	FindIDByExternalService(*entity.ExternalService) (string, bool, error)
	StoreExternalService(service *entity.ExternalService, userID string) error
	FindMetaByEventID(string) ([]*entity.UserMeta, error)
	FindTag() ([]*entity.UserTag, error)
}
