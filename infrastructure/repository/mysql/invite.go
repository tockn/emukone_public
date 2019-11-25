package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type Invite struct {
	ID            int64 `gorm:"primary_key"`
	ArtistUserID  int64
	Artist        User `gorm:"foreignkey:ArtistUserID"`
	EventID       int64
	Event         Event `gorm:"foreignkey:EventID"`
	InviterUserID int64
	Inviter       User `gorm:"foreignkey:InviterUserID"`
	Description   string
	Status        int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Canceled      bool
}

var (
	InviteStatusUnchecked = 0
	InviteStatusApprove   = 1
	InviteStatusDecline   = 2
)

func NewInviteFromDomainAndIDs(ei *entity.Invite, id, artistUserID, eventID, inviterID int64) *Invite {
	return &Invite{
		ID:            id,
		ArtistUserID:  artistUserID,
		EventID:       eventID,
		InviterUserID: inviterID,
		Description:   ei.Description,
		Status:        ei.Status,
	}
}

func (i *Invite) ToDomain(id, artistUserID, eventID, inviterID string) *entity.Invite {
	return &entity.Invite{
		ID:          id,
		ArtistID:    artistUserID,
		Artist:      *i.Artist.ToDomainMeta(artistUserID),
		EventID:     eventID,
		Event:       *i.Event.ToDomain(eventID),
		InviterID:   inviterID,
		Inviter:     *i.Inviter.ToDomainMeta(inviterID),
		Description: i.Description,
		Status:      i.Status,
	}
}
