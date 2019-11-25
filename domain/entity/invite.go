package entity

type Invite struct {
	ID          string   `json:"id"`
	ArtistID    string   `json:"artist_id"`
	Artist      UserMeta `json:"artist"`
	EventID     string   `json:"event_id"`
	Event       Event    `json:"event"`
	InviterID   string   `json:"inviter_id"`
	Inviter     UserMeta `json:"inviter"`
	Description string   `json:"description"`
	Status      int64    `json:"status"`
}

var (
	InviteStatusUnchecked = 0
	InviteStatusApprove   = 1
	InviteStatusDecline   = 2
)
