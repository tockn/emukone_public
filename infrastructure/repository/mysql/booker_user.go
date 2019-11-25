package mysql

import (
	"time"
)

type BookerUser struct {
	UserID          int64  `gorm:"primary_key"`
	Identifier      string `gorm:"default:'booker'"`
	WhyDescription  string
	HowDescription  string
	FreeDescription string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

var BOOKERIDF = "booker"
