package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type Event interface {
	FindByID(string) (*entity.Event, error)
	FindByUserID(string) ([]*entity.Event, error)
	IsYours(id, userID string) bool
	Create(*entity.Event, string) (*entity.Event, error)
	Update(*entity.Event) (*entity.Event, error)
	Delete(*entity.Event) error
}

type event struct {
	Repository repository.Event
}

func NewEvent(repo repository.Event) Event {
	return &event{
		repo,
	}
}

func (s *event) FindByID(id string) (*entity.Event, error) {
	return s.Repository.FindByID(id)
}

func (s *event) FindByUserID(id string) ([]*entity.Event, error) {
	return s.Repository.FindByUserID(id)
}

func (s *event) IsYours(id, userID string) bool {
	_, err := s.Repository.FindByIDAndUserID(id, userID)
	if err != nil {
		return false
	}
	return true
}

func (s *event) Create(du *entity.Event, authorID string) (*entity.Event, error) {
	return s.Repository.Store(du, authorID)
}

func (s *event) Update(du *entity.Event) (*entity.Event, error) {
	return s.Repository.Update(du)
}

func (s *event) Delete(du *entity.Event) error {
	return s.Repository.Delete(du)
}
