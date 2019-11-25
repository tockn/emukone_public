package repository

import "github.com/tockn/emukone/domain/entity"

type Invite interface {
	FindByID(string) (*entity.Invite, error)
	FindByArtistID(string) ([]*entity.Invite, error)
	FindByEventID(string) ([]*entity.Invite, error)
	FindByIDAndArtistID(id, artistID string) (*entity.Invite, error)
	FindByIDAndInviterID(id, inviterID string) (*entity.Invite, error)
	FindByEventIDAndArtistID(eventID, artistID string) (*entity.Invite, error)
	Store(*entity.Invite) (*entity.Invite, error)
	Update(*entity.Invite) (*entity.Invite, error)
	UpdateStatus(*entity.Invite) (*entity.Invite, error)
	Delete(*entity.Invite) error
}
