package handler

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/usecase"
	"net/http"
)

type AuthHandler interface {
	AuthBegin(*gin.Context)
	AuthCallback(*gin.Context)
	AuthLogout(*gin.Context)
}

//
//[WARNING] こいつらはAPIじゃないからな！
type authHandler struct {
	frontEndpoint     string
	userUsecase       usecase.User
	artistUserUsecase usecase.ArtistUser
	bookerUserUsecase usecase.BookerUser
}

func NewAuth(fe string, u usecase.User, ba usecase.ArtistUser, bo usecase.BookerUser) AuthHandler {
	return &authHandler{
		fe,
		u,
		ba,
		bo,
	}
}

// AuthBegin godoc
// @Summary  OAuthでログインする
// @Description これはAPIではないです。認可が成功するとフロントエンドへリダイレクトします。リダイレクト先のパスはqueryで指定できます。
// @Accept  json
// @Produce  json
// @Param provider path string true "OAuthの外部サービス名。twitterとか"
// @Param user-type query string true "登録orログインするユーザーの種類。artist か booker"
// @Param path query string false "リダイレクト先のpathを指定できる。指定しないとルートへ飛ぶ"
// @Success 304
// @Router /auth/{provider} [get]
func (h *authHandler) AuthBegin(c *gin.Context) {
	session := sessions.Default(c)

	// 登録するユーザーのタイプをクッキーに保存しておく
	// ただのログインなら空文字がくる
	userType := c.Query("type")
	session.Set("userType", userType)

	// 認可後にリダイレクトするPATHをクッキーに保存しておく
	path := c.Query("path")
	session.Set("path", path)

	if err := session.Save(); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}

	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h *authHandler) AuthCallback(c *gin.Context) {

	// gothicを使って認可したユーザー情報を取得
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	// AuthBeginでセットしたリダイレクト先のpathを取得
	session := sessions.Default(c)
	path := session.Get("path").(string)

	es := entity.ExternalService{
		Service: user.Provider,
		UID:     user.UserID,
	}

	// 登録済みユーザーかどうかを調べる
	id, registered, err := h.userUsecase.ShowIDByExternalService(&es)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	if registered {

		session.Set("id", id)
		if err := session.Save(); err != nil {
			ErrorResponse(c, ErrBadRequest)
			return
		}

	} else {

		userType, ok := session.Get("userType").(string)
		if !ok {
			ErrorResponse(c, ErrBadRequest)
			return
		}
		session.Delete("userType")

		newID, err := registerUser(h, &user, userType)
		if err != nil {
			ErrorResponse(c, err)
			return
		}

		session.Set("id", newID)
		if err := session.Save(); err != nil {
			ErrorResponse(c, ErrBadRequest)
			return
		}
	}
	c.Redirect(http.StatusTemporaryRedirect, concatPath(h.frontEndpoint, path))
}

func (h *authHandler) AuthLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("id")
	if err := session.Save(); err != nil {
		ErrorResponse(c, ErrBadRequest)
		return
	}
}

func registerUser(h *authHandler, user *goth.User, userType string) (string, error) {
	var newID string
	if userType == entity.ARTISTIDF {
		id, err := h.artistUserUsecase.New(ArtistUserFromGothUser(user))
		if err != nil {
			return "", err
		}
		newID = id
	} else if userType == entity.BOOKERIDF {
		id, err := h.bookerUserUsecase.New(BookerUserFromGothUser(user))
		if err != nil {
			return "", err
		}
		newID = id
	} else {
		return "", ErrBadRequest
	}
	es := ExternalServiceUIDFromGothUser(user)
	if err := h.userUsecase.NewExternalService(es, newID); err != nil {
		return "", err
	}

	return newID, nil
}
