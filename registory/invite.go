package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewInviteHandler() handler.Invite {
	return handler.NewInvite(i.NewInviteUsecase())
}

func (i *interactor) NewInviteUsecase() usecase.Invite {
	return usecase.NewInvite(i.NewInviteService(), i.NewHashIDService(), i.NewEventService())
}

func (i *interactor) NewInviteService() service.Invite {
	return service.NewInvite(i.NewInviteRepository())
}

func (i *interactor) NewInviteRepository() domainRepo.Invite {
	return repository.NewInvite(i.conn, i.NewHashIDRepository())
}
