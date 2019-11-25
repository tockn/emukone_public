package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type UserImage struct {
	ID        int64 `gorm:"primary_key"`
	ImageURL  string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserImageFromDomain(u *entity.UserImage, id int64) *UserImage {
	return &UserImage{
		ImageURL: u.ImageURL,
		UserID:   id,
	}
}

func NewUserImagesFromDomains(dus []*entity.UserImage, id int64) []*UserImage {
	us := make([]*UserImage, len(dus))
	for i, u := range dus {
		us[i] = NewUserImageFromDomain(u, id)
	}
	return us
}

func (i *UserImage) ToDomain() *entity.UserImage {
	return &entity.UserImage{
		ImageURL: i.ImageURL,
	}
}

func UserImagesToDomains(is []*UserImage) []*entity.UserImage {
	dis := make([]*entity.UserImage, len(is))
	for i, im := range is {
		dis[i] = im.ToDomain()
	}
	return dis
}
