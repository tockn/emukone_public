package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type Ichioshi struct {
	ID           int64 `gorm:"primary_key"`
	Title        string
	Service      string
	URL          string
	Embed        string
	ArtistUserID int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewIchioshiFromDomainAndID(u *entity.Ichioshi, id int64) *Ichioshi {
	return &Ichioshi{
		Title:        u.Title,
		Service:      u.Service,
		URL:          u.URL,
		ArtistUserID: id,
		Embed:        u.Embed,
	}
}

func NewIchioshisFromDomainAndID(dus []*entity.Ichioshi, id int64) []*Ichioshi {
	us := make([]*Ichioshi, len(dus))
	for i, u := range dus {
		us[i] = NewIchioshiFromDomainAndID(u, id)
	}
	return us
}

func (i *Ichioshi) ToDomain() *entity.Ichioshi {
	return &entity.Ichioshi{
		Title:   i.Title,
		Service: i.Service,
		URL:     i.URL,
		Embed:   i.Embed,
	}
}

func IchioshisToDomains(is []*Ichioshi) []*entity.Ichioshi {
	dis := make([]*entity.Ichioshi, len(is))
	for i, ic := range is {
		dis[i] = ic.ToDomain()
	}
	return dis
}
