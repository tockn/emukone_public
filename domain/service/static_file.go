package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type StaticFile interface {
	UploadImage(*entity.Image) (*entity.Image, error)
}

type staticFile struct {
	repo repository.StaticFile
}

func NewStaticFile(repo repository.StaticFile) StaticFile {
	return &staticFile{
		repo,
	}
}

func (s *staticFile) UploadImage(img *entity.Image) (*entity.Image, error) {
	return s.repo.StoreImage(img)
}
