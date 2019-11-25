package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type Location struct {
	ID        int64 `gorm:"primary_key"`
	Name      string
	Users     []*User `gorm:"many2many:locations_related_users;"`
	CreatedAt time.Time
}

func (l *Location) ToDomain() *entity.Location {
	return &entity.Location{
		Name: l.Name,
	}
}

func LocationsToDomains(ls []*Location) []*entity.Location {
	dls := make([]*entity.Location, len(ls))
	for i, l := range ls {
		dls[i] = l.ToDomain()
	}
	return dls
}

func NewLocationFromDomain(l *entity.Location) *Location {
	return &Location{
		Name: l.Name,
	}
}

func NewLocationsFromDomains(dl []*entity.Location) []*Location {
	ls := make([]*Location, len(dl))
	for i, l := range dl {
		ls[i] = NewLocationFromDomain(l)
	}
	return ls
}

type LocationsRelatedUsers struct {
	ID         int64 `gorm:"primary_key"`
	UserID     int64
	LocationID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
