package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/tockn/emukone/domain/entity"
	"github.com/tockn/emukone/infrastructure/repository"
	"github.com/tockn/emukone/usecase"
	"log"
	"net/http"
	"strings"
)

var (
	ErrBadRequest = errors.New("Bad Request")
)

func ErrorResponse(c *gin.Context, err error) {
	log.Println(err)
	if err == gorm.ErrRecordNotFound || err == sql.ErrNoRows || err == repository.ErrDecode {
		c.JSON(http.StatusNotFound, "Not Found")
	} else if err == ErrBadRequest {
		c.JSON(http.StatusBadRequest, err.Error())
	} else if err == usecase.ErrUnAuthorized {
		c.JSON(http.StatusUnauthorized, err)
	} else if err == usecase.ErrAlreadyInvite {
		c.JSON(http.StatusConflict, err)
	} else {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

func retrieveID(c *gin.Context) string {
	s := sessions.Default(c)
	id, _ := s.Get("id").(string)
	return id
}

// concatPathは、スラッシュの関係が変でもURLパスをいい感じに結合する
// 例えば http://localhost/ と /hoge でも
// http://localhost と hoge でも
// http://localhost/ と hoge でも
// 全部 http://localhost/hoge にしてくれる
func concatPath(path1, path2 string) string {
	if len(path1) == 0 || len(path2) == 0 {
		return path1 + path2
	}
	path1End := string(path1[len(path1)-1])
	path2Head := string(path2[0])
	if path1End == "/" && path2Head == "/" {
		return path1 + string([]rune(path2)[1:])
	} else if path1End != "/" && path2Head != "/" {
		return path1 + "/" + path2
	}
	return path1 + path2
}

func ExternalServiceUIDFromGothUser(gu *goth.User) *entity.ExternalService {
	return &entity.ExternalService{
		Service: gu.Provider,
		UID:     gu.UserID,
	}
}

func UserFromGothUser(gu *goth.User) *entity.User {
	if gu.Provider == "twitter" {
		gu.AvatarURL = strings.Replace(gu.AvatarURL, "_normal", "", 1)
	}

	return &entity.User{
		Name:    gu.Name,
		IconURL: gu.AvatarURL,
	}
}

func ArtistUserFromGothUser(gu *goth.User) *entity.ArtistUser {
	if gu.Provider == "twitter" {
		gu.AvatarURL = strings.Replace(gu.AvatarURL, "_normal", "", 1)
	}

	return &entity.ArtistUser{
		Name:            gu.Name,
		IconURL:         gu.AvatarURL,
		FreeDescription: gu.Description,
	}
}

func BookerUserFromGothUser(gu *goth.User) *entity.BookerUser {
	if gu.Provider == "twitter" {
		gu.AvatarURL = strings.Replace(gu.AvatarURL, "_normal", "", 1)
	}

	return &entity.BookerUser{
		Name:            gu.Name,
		IconURL:         gu.AvatarURL,
		FreeDescription: gu.Description,
	}
}
