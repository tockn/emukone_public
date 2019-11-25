package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type Invite interface {
	GetInvite(*gin.Context)
	GetInvitesByArtistID(*gin.Context)
	GetInvitesByEventID(*gin.Context)
	PostInvite(*gin.Context)
	PutInvite(*gin.Context)
	PutInviteStatus(*gin.Context)
	DeleteInvite(*gin.Context)
}

type invite struct {
	usecase usecase.Invite
}

func NewInvite(u usecase.Invite) Invite {
	return &invite{
		u,
	}
}

// GetInvite godoc
// @Summary  Get invite
// @Description
// @Accept  json
// @Produce  json
// @Param inviteID path string true "invite id"
// @Success 200 {object} entity.Invite
// @Router /invites/{inviteID} [get]
func (h *invite) GetInvite(c *gin.Context) {
	id := c.Param("inviteID")
	userID := retrieveID(c)
	di, err := h.usecase.Show(id, userID)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, di)
}

// GetInvitesByArtistID godoc
// @Summary  バンドIDからスカウトを取得する
// @Description
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Invite
// @Router /invites [get]
func (h *invite) GetInvitesByArtistID(c *gin.Context) {
	id := retrieveID(c)
	di, err := h.usecase.ShowByUserID(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, di)
}

// GetInvitesByEventID godoc
// @Summary  イベントIDからスカウトを取得する
// @Description
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Invite
// @Router /events/{eventID}/invites [get]
func (h *invite) GetInvitesByEventID(c *gin.Context) {
	eventID := c.Param("eventID")
	userID := retrieveID(c)
	di, err := h.usecase.ShowByEventID(eventID, userID)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, di)
}

// PostInvite godoc
// @Summary  スカウトする
// @Description
// @Accept  json
// @Produce  json
// @Param invite body entity.Invite true "invite"
// @Success 201 {object} entity.Invite
// @Router /invites [post]
func (h *invite) PostInvite(c *gin.Context) {
	var ei entity.Invite
	if err := c.BindJSON(&ei); err != nil {
		ErrorResponse(c, err)
		return
	}

	ei.InviterID = retrieveID(c)

	nei, err := h.usecase.New(&ei)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, nei)
}

// PostInvite godoc
// @Summary  スカウトの内容（説明とか）を編集する
// @Description
// @Accept  json
// @Produce  json
// @Param invite body entity.Invite true "edited invite"
// @Success 204 {object} entity.Invite
// @Router /invites [put]
func (h *invite) PutInvite(c *gin.Context) {
	var di entity.Invite
	if err := c.BindJSON(&di); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	di.InviterID = retrieveID(c)

	ndu, err := h.usecase.Edit(&di)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, ndu)
}

// PutInviteStatus godoc
// @Summary  スカウトの承認状態を編集する
// @Description
// @Accept  json
// @Produce  json
// @Param inviteID path string true "承認非承認するスカウトID"
// @Param invite body entity.Invite true "edited invite"
// @Success 200 {object} entity.Invite
// @Router /invites/{inviteID}/status [put]
func (h *invite) PutInviteStatus(c *gin.Context) {
	var ei entity.Invite
	if err := c.BindJSON(&ei); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	ei.ID = c.Param("inviteID")
	ei.ArtistID = retrieveID(c)

	nei, err := h.usecase.EditStatus(&ei)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nei)
}

// DeleteInvite godoc
// @Summary  delete invite
// @Description
// @Accept  json
// @Produce  json
// @Param invite body entity.Invite true "invite"
// @Success 204
// @Router /invites [delete]
func (h *invite) DeleteInvite(c *gin.Context) {
	var di entity.Invite
	if err := c.BindJSON(&di); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	di.InviterID = retrieveID(c)

	if err := h.usecase.Delete(&di); err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
