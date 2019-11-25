package registory

import (
	domainRepo "github.com/tockn/emukone/domain/repository"
	"github.com/tockn/emukone/domain/service"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/usecase"
)

func (i *interactor) NewStaticFileHandler() handler.StaticFile {
	return handler.NewStaticImage(i.NewStaticFileUsecase())
}

func (i *interactor) NewStaticFileUsecase() usecase.StaticFile {
	return usecase.NewStaticFile(i.NewStaticFileService())
}

func (i *interactor) NewStaticFileService() service.StaticFile {
	return service.NewStaticFile(i.NewStaticFileRepository())
}

func (i *interactor) NewStaticFileRepository() domainRepo.StaticFile {
	return repository.NewStaticFile(i.awsSession)
}
