package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewUserHandler() handler.User {
	return handler.NewUser(i.NewUserUsecase())
}

func (i *interactor) NewUserUsecase() usecase.User {
	return usecase.NewUser(i.NewUserService(), i.NewHashIDService())
}

func (i *interactor) NewUserService() service.User {
	return service.NewUser(i.NewUserRepository())
}

func (i *interactor) NewUserRepository() domainRepo.User {
	return repository.NewUser(i.conn, repository.NewHashID(i.hash))
}
