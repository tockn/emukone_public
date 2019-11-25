package repository

import "github.com/tockn/emukone/domain/entity"

type StaticFile interface {
	StoreImage(*entity.Image) (*entity.Image, error)
}
