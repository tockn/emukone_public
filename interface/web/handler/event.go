package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type Event interface {
	GetEvent(c *gin.Context)
	GetEventsByUserID(c *gin.Context)
	PostEvent(c *gin.Context)
	PutEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
}

type event struct {
	usecase usecase.Event
}

func NewEvent(e usecase.Event) Event {
	return &event{
		e,
	}
}

// @Summary  get event
// @Description
// @Accept  json
// @Produce  json
// @Param eventID path string true "event id"
// @Success 200 {object} entity.Event
// @Router /events/{eventID} [get]
func (h *event) GetEvent(c *gin.Context) {
	id := c.Param("eventID")
	du, err := h.usecase.Show(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, du)
}

// GetEventsByUserID godoc
// @Summary  ユーザーIDからイベントを取得
// @Description
// @Accept  json
// @Produce  json
// @Param userID path string true "userID"
// @Success 200 {object} entity.Event
// @Router /users/{userID}/events [get]

func (h *event) GetEventsByUserID(c *gin.Context) {
	id := c.Param("userID")
	des, err := h.usecase.ShowByUserID(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, des)
}

// PostEvent godoc
// @Summary  Post event
// @Description
// @Accept  json
// @Produce  json
// @Param event body entity.Event true "event"
// @Success 201 {object} entity.Event
// @Router /events [post]
func (h *event) PostEvent(c *gin.Context) {
	var du entity.Event
	if err := c.BindJSON(&du); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	authorID := retrieveID(c)

	ndu, err := h.usecase.New(&du, authorID)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, ndu)
}

// PutEvent godoc
// @Summary  Edit event
// @Description
// @Accept  json
// @Produce  json
// @Param event body entity.Event true "event"
// @Success 204 {object} entity.Event
// @Router /events [put]
func (h *event) PutEvent(c *gin.Context) {
	var du entity.Event
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

// DeleteEvent godoc
// @Summary  delete event
// @Description
// @Accept  json
// @Produce  json
// @Param event body entity.Event true "event"
// @Success 204
// @Router /events [delete]
func (h *event) DeleteEvent(c *gin.Context) {
	var du entity.Event
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
