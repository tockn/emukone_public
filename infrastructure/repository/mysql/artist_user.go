package mysql

import (
	"time"
)

type ArtistUser struct {
	UserID          int64  `gorm:"primary_key"`
	Identifier      string `gorm:"default:'artist'"`
	WhyDescription  string
	HowDescription  string
	FreeDescription string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Invites []*Invite `gorm:"foreignkey:ArtistUserID"`
}

var ARTISTIDF = "artist"
