package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type BookerUser interface {
	GetBookerUser(c *gin.Context)
	PostBookerUser(c *gin.Context)
	PutBookerUser(c *gin.Context)
	DeleteBookerUser(c *gin.Context)
}

type bookerUser struct {
	usecase usecase.BookerUser
}

func NewBookerUser(c usecase.BookerUser) BookerUser {
	return &bookerUser{
		c,
	}
}

// GetBookerUser godoc
// @Summary  Get Booker User
// @Description
// @Accept  json
// @Produce  json
// @Param bookerUserID path string true "booker user id"
// @Success 200 {object} entity.BookerUser
// @Router /bookers/{bookerUserID} [get]
func (h *bookerUser) GetBookerUser(c *gin.Context) {
	id := c.Param("bookerID")
	du, err := h.usecase.Show(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, du)
}

func (h *bookerUser) PostBookerUser(c *gin.Context) {
	var du entity.BookerUser
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

// PutBookerUser godoc
// @Summary  edit booker user
// @Description
// @Accept  json
// @Produce  json
// @Param bookerUser body entity.BookerUser true "edited booker user json"
// @Success 204 {object} entity.BookerUser
// @Router /bookers [put]
func (h *bookerUser) PutBookerUser(c *gin.Context) {
	var du entity.BookerUser
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

// DeleteBookerUser godoc
// @Summary  delete booker user
// @Description
// @Accept  json
// @Produce  json
// @Param bookerUser body entity.BookerUser true "delete booker user"
// @Success 204
// @Router /bookers [delete]
func (h *bookerUser) DeleteBookerUser(c *gin.Context) {
	var du entity.BookerUser
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
