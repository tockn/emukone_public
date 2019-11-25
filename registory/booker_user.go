package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewBookerUserHandler() handler.BookerUser {
	return handler.NewBookerUser(i.NewBookerUserController())
}

func (i *interactor) NewBookerUserController() usecase.BookerUser {
	return usecase.NewBookerUser(i.NewBookerUserService(), i.NewHashIDService())
}

func (i *interactor) NewBookerUserService() service.BookerUser {
	return service.NewBookerUser(i.NewBookerUserRepository())
}

func (i *interactor) NewBookerUserRepository() domainRepo.BookerUser {
	return repository.NewBookerUser(i.conn, repository.NewHashID(i.hash))
}
