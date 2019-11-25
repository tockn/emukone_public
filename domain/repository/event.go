package repository

import "github.com/tockn/emukone/domain/entity"

type Event interface {
	FindByID(string) (*entity.Event, error)
	FindByUserID(string) ([]*entity.Event, error)
	FindByIDAndUserID(id, userID string) (*entity.Event, error)
	Store(*entity.Event, string) (*entity.Event, error)
	Update(*entity.Event) (*entity.Event, error)
	Delete(*entity.Event) error
}
