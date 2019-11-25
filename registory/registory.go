package registory

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jinzhu/gorm"
	"github.com/speps/go-hashids"
	"github.com/tockn/emukone/interface/web/handler"
)

type Interactor interface {
	NewAppHandler() handler.Application
}

type interactor struct {
	conn          *gorm.DB
	hash          *hashids.HashID
	frontEndpoint string
	awsSession    *session.Session
}

func NewInteractor(conn *gorm.DB, hash *hashids.HashID, fe string, awsSession *session.Session) Interactor {
	return &interactor{
		conn:          conn,
		hash:          hash,
		frontEndpoint: fe,
		awsSession:    awsSession,
	}
}

func (i *interactor) NewAppHandler() handler.Application {
	return struct {
		handler.User
		handler.BookerUser
		handler.ArtistUser
		handler.AuthHandler
		handler.Event
		handler.Invite
		handler.Tag
		handler.StaticFile
	}{
		User:        i.NewUserHandler(),
		BookerUser:  i.NewBookerUserHandler(),
		ArtistUser:  i.NewArtistUserHandler(),
		AuthHandler: i.NewAuthHandler(),
		Event:       i.NewEventHandler(),
		Invite:      i.NewInviteHandler(),
		Tag:         i.NewTagHandler(),
		StaticFile:  i.NewStaticFileHandler(),
	}
}
