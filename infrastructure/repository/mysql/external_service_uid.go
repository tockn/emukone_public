package mysql

import (
	"github.com/tockn/emukone/domain/entity"
	"time"
)

type ExternalServiceUID struct {
	ID        int64 `gorm:"primary_key"`
	Service   string
	UID       string
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewExternalServiceUIDFromDomainAndUserID(d *entity.ExternalService, userID int64) *ExternalServiceUID {
	return &ExternalServiceUID{
		Service: d.Service,
		UID:     d.UID,
		UserID:  userID,
	}
}

func (e *ExternalServiceUID) ToDomain() *entity.ExternalService {
	return &entity.ExternalService{
		Service: e.Service,
		UID:     e.UID,
	}
}
