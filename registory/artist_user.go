package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewArtistUserHandler() handler.ArtistUser {
	return handler.NewArtistUser(i.NewArtistUserUsecase())
}

func (i *interactor) NewArtistUserUsecase() usecase.ArtistUser {
	return usecase.NewArtistUser(i.NewArtistUserService(), i.NewHashIDService())
}

func (i *interactor) NewArtistUserService() service.ArtistUser {
	return service.NewArtistUser(i.NewArtistUserRepository())
}

func (i *interactor) NewArtistUserRepository() domainRepo.ArtistUser {
	return repository.NewArtistUser(i.conn, repository.NewHashID(i.hash))
}
