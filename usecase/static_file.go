package usecase

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type StaticFile interface {
	NewImage(*entity.Image) (*entity.Image, error)
}

type staticFile struct {
	staticFileService service.StaticFile
}

func NewStaticFile(sf service.StaticFile) StaticFile {
	return &staticFile{
		sf,
	}
}

func (u *staticFile) NewImage(img *entity.Image) (*entity.Image, error) {
	return u.staticFileService.UploadImage(img)
}
