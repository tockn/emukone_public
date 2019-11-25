package registory

import "github.com/tockn/emukone/interface/web/handler"

func (i *interactor) NewTagHandler() handler.Tag {
	return handler.NewTag(i.NewUserUsecase())
}
