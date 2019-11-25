package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type ArtistMember struct {
	ID           int64 `gorm:"primary_key"`
	Name         string
	Part         string
	ArtistUserID int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *ArtistMember) ToDoamin() *entity.ArtistMember {
	return &entity.ArtistMember{
		Name: m.Name,
		Part: m.Part,
	}
}

func ArtistMembersToDomains(bs []*ArtistMember) []*entity.ArtistMember {
	dbs := make([]*entity.ArtistMember, len(bs))
	for i, b := range bs {
		dbs[i] = b.ToDoamin()
	}
	return dbs
}

func NewArtistMemberFromDomainAndID(u *entity.ArtistMember, id int64) *ArtistMember {
	return &ArtistMember{
		Name:         u.Name,
		Part:         u.Part,
		ArtistUserID: id,
	}
}

func NewArtistMembersFromDomainsAndID(dus []*entity.ArtistMember, id int64) []*ArtistMember {
	us := make([]*ArtistMember, len(dus))
	for i, u := range dus {
		us[i] = NewArtistMemberFromDomainAndID(u, id)
	}
	return us
}
