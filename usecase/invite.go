package usecase

import (
	"errors"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/domain/service"
)

type Invite interface {
	Show(string, string) (*entity.Invite, error)
	ShowByUserID(string) ([]*entity.Invite, error)
	ShowByEventID(eventID, userID string) ([]*entity.Invite, error)
	New(*entity.Invite) (*entity.Invite, error)
	Edit(*entity.Invite) (*entity.Invite, error)
	EditStatus(*entity.Invite) (*entity.Invite, error)
	Delete(*entity.Invite) error
}

var (
	ErrAlreadyInvite = errors.New("already invite")
)

type invite struct {
	inviteService service.Invite
	hashIDService service.HashID
	eventService  service.Event
}

func NewInvite(is service.Invite, hs service.HashID, es service.Event) Invite {
	return &invite{
		inviteService: is,
		hashIDService: hs,
		eventService:  es,
	}
}

func (u *invite) Show(id, userID string) (*entity.Invite, error) {
	if !(u.inviteService.IsInvite(id, userID) || u.inviteService.IsInvited(id, userID)) {
		return nil, ErrUnAuthorized
	}

	return u.inviteService.FindByID(id)
}

func (u *invite) ShowByUserID(id string) ([]*entity.Invite, error) {
	return u.inviteService.FindByUserID(id)
}

func (u *invite) ShowByEventID(eventID, userID string) ([]*entity.Invite, error) {
	if !u.eventService.IsYours(eventID, userID) {
		return nil, ErrUnAuthorized
	}
	return u.inviteService.FindByEventID(eventID)
}

func (u *invite) New(ei *entity.Invite) (*entity.Invite, error) {

	if !u.eventService.IsYours(ei.EventID, ei.InviterID) {
		return nil, ErrUnAuthorized
	}

	if u.inviteService.CheckExistByEventIDAndArtistID(ei.EventID, ei.ArtistID) {
		return nil, ErrAlreadyInvite
	}

	nei, err := u.inviteService.Create(ei)
	if err != nil {
		return nil, err
	}

	// TODO バンドユーザーに通知する

	return nei, nil
}

func (u *invite) Edit(ei *entity.Invite) (*entity.Invite, error) {

	if !u.inviteService.IsInvite(ei.ID, ei.InviterID) {
		return nil, ErrUnAuthorized
	}

	return u.inviteService.Update(ei)
}

func (u *invite) EditStatus(ei *entity.Invite) (*entity.Invite, error) {
	if !u.inviteService.IsInvited(ei.ID, ei.ArtistID) {
		return nil, ErrUnAuthorized
	}
	return u.inviteService.UpdateStatus(ei)
}

func (u *invite) Delete(ei *entity.Invite) error {
	return u.inviteService.Delete(ei)
}
