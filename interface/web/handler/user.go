package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type User interface {
	GetUserMeta(c *gin.Context)
	GetMe(c *gin.Context)
	GetUserMetaByEventID(c *gin.Context)
}

type user struct {
	usecase usecase.User
}

func NewUser(c usecase.User) User {
	return &user{
		c,
	}
}

// GetUserMeta godoc
// @Summary  ユーザーのメタ情報（概略）を取得する
// @Description
// @Accept  json
// @Produce  json
// @Param userID path string true "user id"
// @Success 200 {object} entity.UserMeta
// @Router /users/{userID}/meta [get]
func (h *user) GetUserMeta(c *gin.Context) {
	id := c.Param("userID")
	du, err := h.usecase.ShowMeta(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, du)
}

// @Summary  なんのユーザーとしてログインしているかを調べる
// @Description なんのユーザーとしてログインしているかを調べる。ログインしていたらUserMetaが帰ってくるが、ログインしていなければエラー。
// @Produce  json
// @Success 200 {object} entity.UserMeta
// @Router /me [get]
func (h *user) GetMe(c *gin.Context) {
	id := retrieveID(c)
	du, err := h.usecase.ShowMeta(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, du)
}

func (h *user) GetUserMetaByEventID(c *gin.Context) {
	id := c.Param("eventID")
	dus, err := h.usecase.ShowMetaByEventID(id)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, dus)
}
