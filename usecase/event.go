package usecase

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type Event interface {
	Show(string) (*entity.Event, error)
	ShowByUserID(string) ([]*entity.Event, error)
	New(*entity.Event, string) (string, error)
	Edit(*entity.Event) (*entity.Event, error)
	Delete(*entity.Event) error
}

type event struct {
	eventService  service.Event
	hashIDService service.HashID
}

func NewEvent(e service.Event, h service.HashID) Event {
	return &event{
		e,
		h,
	}
}

func (u *event) Show(id string) (*entity.Event, error) {
	return u.eventService.FindByID(id)
}

func (u *event) ShowByUserID(id string) ([]*entity.Event, error) {
	return u.eventService.FindByUserID(id)
}

func (u *event) New(du *entity.Event, authorID string) (string, error) {
	nu, err := u.eventService.Create(du, authorID)
	if err != nil {
		return "", err
	}
	return nu.ID, nil
}

func (u *event) Edit(du *entity.Event) (*entity.Event, error) {
	return u.eventService.Update(du)
}

func (u *event) Delete(du *entity.Event) error {
	return u.eventService.Delete(du)
}
