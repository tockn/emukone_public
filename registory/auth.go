package registory

import "github.com/tockn/emukone/interface/web/handler"

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuth(i.frontEndpoint, i.NewUserUsecase(), i.NewArtistUserUsecase(), i.NewBookerUserController())
}
