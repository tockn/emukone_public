package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type Tag interface {
	GetAllTags(*gin.Context)
}

type tag struct {
	userUsecase usecase.User
}

func NewTag(u usecase.User) Tag {
	return &tag{
		u,
	}
}

func (h *tag) GetAllTags(c *gin.Context) {
	tags, err := h.userUsecase.ShowAllTags()
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, tags)
}
