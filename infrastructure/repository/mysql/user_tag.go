package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type UserTag struct {
	ID        int64 `gorm:"primary_key"`
	Name      string
	Users     []*User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserTagFromDomain(t *entity.UserTag) *UserTag {
	return &UserTag{
		Name: t.Name,
	}
}

func NewUserTagsFromDomains(dt []*entity.UserTag) []*UserTag {
	ts := make([]*UserTag, len(dt))
	for i, t := range dt {
		ts[i] = NewUserTagFromDomain(t)
	}
	return ts
}

func (t *UserTag) ToDomain() *entity.UserTag {
	return &entity.UserTag{
		Name: t.Name,
	}
}

func UserTagsToDomains(ts []*UserTag) []*entity.UserTag {
	dts := make([]*entity.UserTag, len(ts))
	for i, t := range ts {
		dts[i] = t.ToDomain()
	}
	return dts
}

type UserTagsRelatedUsers struct {
	ID        int64 `gorm:"primary_key"`
	UserID    int64
	UserTagID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
