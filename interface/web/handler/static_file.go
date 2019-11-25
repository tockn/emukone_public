package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type StaticFile interface {
	PostImage(*gin.Context)
}

type staticImage struct {
	usecase usecase.StaticFile
}

func NewStaticImage(usecase usecase.StaticFile) StaticFile {
	return &staticImage{
		usecase,
	}
}

func (h *staticImage) PostImage(c *gin.Context) {
	var img *entity.Image
	if err := c.BindJSON(&img); err != nil {
		ErrorResponse(c, err)
		return
	}
	newImg, err := h.usecase.NewImage(img)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, newImg)
}
