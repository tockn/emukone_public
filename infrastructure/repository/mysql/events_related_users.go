package mysql

import "time"

type EventsRelatedUsers struct {
	ID        int64 `gorm:"primary_key"`
	EventID   int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEventsRelatedUsers(eventID, userID int64) *EventsRelatedUsers {
	return &EventsRelatedUsers{
		EventID: eventID,
		UserID:  userID,
	}
}
