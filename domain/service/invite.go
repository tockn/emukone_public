package service

import (
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/repository"
)

type Invite interface {
	FindByID(string) (*entity.Invite, error)
	FindByUserID(string) ([]*entity.Invite, error)
	FindByEventID(string) ([]*entity.Invite, error)
	IsInvited(inviteID, artistID string) bool
	IsInvite(inviteID, inviterID string) bool
	CheckExistByEventIDAndArtistID(eventID, artistID string) bool
	Create(*entity.Invite) (*entity.Invite, error)
	Update(*entity.Invite) (*entity.Invite, error)
	UpdateStatus(*entity.Invite) (*entity.Invite, error)
	Delete(*entity.Invite) error
}

type invite struct {
	Repo repository.Invite
}

func NewInvite(repo repository.Invite) Invite {
	return &invite{
		repo,
	}
}

func (s *invite) FindByID(id string) (*entity.Invite, error) {
	return s.Repo.FindByID(id)
}

func (s *invite) FindByUserID(id string) ([]*entity.Invite, error) {
	return s.Repo.FindByArtistID(id)
}

// TODO
func (s *invite) FindByInviterID(id string) ([]*entity.Invite, error) {
	return nil, nil
}

func (s *invite) FindByEventID(id string) ([]*entity.Invite, error) {
	return s.Repo.FindByEventID(id)
}

// Is this the invite that this artist is invited to?
func (s *invite) IsInvited(id, artistID string) bool {
	_, err := s.Repo.FindByIDAndArtistID(id, artistID)
	if err != nil {
		return false
	}
	return true
}

func (s *invite) IsInvite(id, inviterID string) bool {
	_, err := s.Repo.FindByIDAndInviterID(id, inviterID)
	if err != nil {
		return false
	}
	return true
}

func (s *invite) CheckExistByEventIDAndArtistID(eventID, artistID string) bool {
	_, err := s.Repo.FindByEventIDAndArtistID(eventID, artistID)
	if err != nil {
		return false
	}
	return true
}

func (s *invite) Create(ei *entity.Invite) (*entity.Invite, error) {
	return s.Repo.Store(ei)
}

func (s *invite) Update(du *entity.Invite) (*entity.Invite, error) {
	return s.Repo.Update(du)
}

func (s *invite) UpdateStatus(du *entity.Invite) (*entity.Invite, error) {
	return s.Repo.UpdateStatus(du)
}

func (s *invite) Delete(du *entity.Invite) error {
	return s.Repo.Delete(du)
}
