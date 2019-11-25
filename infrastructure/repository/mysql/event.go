package mysql

import (
	"errors"
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type Event struct {
	ID              int64 `gorm:"primary_key"`
	Name            string
	OpenTime        time.Time
	StartTime       time.Time
	MetaDescription string
	WhyDescription  string
	FreeDescription string
	Place           string
	TicketPrice     string
	HeaderImageURL  string
	EventImages     []*EventImage
	Canceled        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Users []*User `gorm:"many2many:events_related_users"`

	Invites []*Invite `gorm:"foreignkey:EventID"`
}

func NewEventFromDomainAndID(de *entity.Event, id int64) *Event {
	return &Event{
		ID:              id,
		Name:            de.Name,
		OpenTime:        de.OpenTime,
		StartTime:       de.StartTime,
		MetaDescription: de.MetaDescription,
		WhyDescription:  de.WhyDescription,
		FreeDescription: de.FreeDescription,
		Place:           de.Place,
		TicketPrice:     de.TicketPrice,
		HeaderImageURL:  de.HeaderImageURL,
		EventImages:     NewEventImagesFromDomains(de.EventImages, id),
		Canceled:        false,
	}
}

func (e *Event) ToDomain(idStr string) *entity.Event {
	return &entity.Event{
		ID:              idStr,
		Name:            e.Name,
		MetaDescription: e.MetaDescription,
		WhyDescription:  e.WhyDescription,
		FreeDescription: e.FreeDescription,
		Place:           e.Place,
		TicketPrice:     e.TicketPrice,
		OpenTime:        e.OpenTime,
		StartTime:       e.StartTime,
		HeaderImageURL:  e.HeaderImageURL,
		EventImages:     EventImagesToDomains(e.EventImages),
	}
}

func EventToDomains(es []*Event, ids []string) ([]*entity.Event, error) {
	if len(es) != len(ids) {
		return nil, errors.New("Event len not equals to ids len")
	}
	ees := make([]*entity.Event, len(es))
	for i, e := range es {
		ees[i] = e.ToDomain(ids[i])
	}
	return ees, nil
}
