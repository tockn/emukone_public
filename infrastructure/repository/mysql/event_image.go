package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type EventImage struct {
	ID        int64 `gorm:"primary_key"`
	ImageURL  string
	EventID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEventImageFromDomain(u *entity.EventImage, id int64) *EventImage {
	return &EventImage{
		ImageURL: u.ImageURL,
		EventID:  id,
	}
}

func NewEventImagesFromDomains(dus []*entity.EventImage, id int64) []*EventImage {
	us := make([]*EventImage, len(dus))
	for i, u := range dus {
		us[i] = NewEventImageFromDomain(u, id)
	}
	return us
}

func (i *EventImage) ToDomain() *entity.EventImage {
	return &entity.EventImage{
		ImageURL: i.ImageURL,
	}
}

func EventImagesToDomains(is []*EventImage) []*entity.EventImage {
	dis := make([]*entity.EventImage, len(is))
	for i, im := range is {
		dis[i] = im.ToDomain()
	}
	return dis
}
