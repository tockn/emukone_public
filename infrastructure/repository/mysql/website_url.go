package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type WebsiteURL struct {
	ID        int64 `gorm:"primary_key"`
	Type      string
	URL       string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWebsiteURLFromDomain(u *entity.WebsiteURL, id int64) *WebsiteURL {
	return &WebsiteURL{
		Type:   u.Type,
		URL:    u.URL,
		UserID: id,
	}
}

func NewWebsiteURLsFromDomains(dus []*entity.WebsiteURL, id int64) []*WebsiteURL {
	us := make([]*WebsiteURL, len(dus))
	for i, u := range dus {
		us[i] = NewWebsiteURLFromDomain(u, id)
	}
	return us
}

func (u *WebsiteURL) ToDomain() *entity.WebsiteURL {
	return &entity.WebsiteURL{
		Type: u.Type,
		URL:  u.URL,
	}
}

func WebsiteURLsToDomains(us []*WebsiteURL) []*entity.WebsiteURL {
	dus := make([]*entity.WebsiteURL, len(us))
	for i, u := range us {
		dus[i] = u.ToDomain()
	}
	return dus
}
