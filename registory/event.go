package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewEventHandler() handler.Event {
	return handler.NewEvent(i.NewEventUsecase())
}

func (i *interactor) NewEventUsecase() usecase.Event {
	return usecase.NewEvent(i.NewEventService(), i.NewHashIDService())
}

func (i *interactor) NewEventService() service.Event {
	return service.NewEvent(i.NewEventRepository())
}

func (i *interactor) NewEventRepository() domainRepo.Event {
	return repository.NewEvent(i.conn, i.NewHashIDRepository())
}
