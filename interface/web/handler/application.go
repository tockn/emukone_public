package handler

type Application interface {
	User
	BookerUser
	ArtistUser
	AuthHandler
	Event
	Invite
	Tag
	StaticFile
}
