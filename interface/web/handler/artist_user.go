package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type ArtistUser interface {
	GetArtistUser(c *gin.Context)
	PostArtistUser(c *gin.Context)
	PutArtistUser(c *gin.Context)
	DeleteArtistUser(c *gin.Context)
}

type artistUser struct {
	usecase usecase.ArtistUser
}

func NewArtistUser(c usecase.ArtistUser) ArtistUser {
	return &artistUser{
		c,
	}
}

// GetArtistUser godoc
//// @Summary  Get artist user
//// @Accept  json
//// @Produce  json
//// @Param artistUserID path string true "artist user id"
//// @Success 200 {object} entity.ArtistUser
//// @Router /artists/{artistUserID} [get]
func (h *artistUser) GetArtistUser(c *gin.Context) {
	id := c.Param("artistID")
	du, err := h.usecase.Show(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, du)
}

func (h *artistUser) PostArtistUser(c *gin.Context) {
	var du entity.ArtistUser
	if err := c.BindJSON(&du); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}
	ndu, err := h.usecase.New(&du)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, ndu)
}

// PutArtistUser godoc
// @Summary  edit Artist user
// @Description
// @Accept  json
// @Produce  json
// @Param  artistUser body entity.ArtistUser true "edited artist user"
// @Success 204 {object} entity.ArtistUser
// @Router /artists [put]
func (h *artistUser) PutArtistUser(c *gin.Context) {
	var du entity.ArtistUser
	if err := c.BindJSON(&du); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	du.ID = retrieveID(c)

	ndu, err := h.usecase.Edit(&du)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, ndu)
}

// DeleteArtistUser godoc
// @Summary  Delete Artist User
// @Description
// @Accept  json
// @Produce  json
// @Param artistUser body entity.ArtistUser true "delete artist user"
// @Success 201
// @Router /artists [delete]
func (h *artistUser) DeleteArtistUser(c *gin.Context) {
	var du entity.ArtistUser
	if err := c.BindJSON(&du); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	du.ID = retrieveID(c)

	if err := h.usecase.Delete(&du); err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
