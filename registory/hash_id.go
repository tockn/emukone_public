package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
)

func (i *interactor) NewHashIDService() service.HashID {
	return service.NewHashID(i.NewHashIDRepository())
}

func (i *interactor) NewHashIDRepository() domainRepo.HashID {
	return repository.NewHashID(i.hash)
}
