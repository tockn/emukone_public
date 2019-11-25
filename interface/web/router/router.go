package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tockn/emukone/interface/web/handler"
	"github.com/tockn/emukone/interface/web/middleware"
)

func NewRouter(r *gin.Engine, h handler.Application, store sessions.Store) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:3000", "http://192.168.0.140:3000"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		AllowMethods:     []string{"PUT", "DELETE"},
	}))

	r.Use(sessions.Sessions("emukone", store))

	v1 := r.Group("/v1")
	{

		v1.GET("/me", middleware.AuthRequired(), h.GetMe)

		users := v1.Group("/users")
		{
			users.GET(":userID/meta", h.GetUserMeta)
			users.GET(":userID/events", h.GetEventsByUserID)
		}

		bookers := v1.Group("/bookers")
		{
			bookers.GET(":bookerID", h.GetBookerUser)
			bookers.POST("", middleware.AuthRequired(), h.PostBookerUser)
			bookers.PUT("", middleware.AuthRequired(), h.PutBookerUser)
			bookers.DELETE("", middleware.AuthRequired(), h.DeleteBookerUser)
		}
		artists := v1.Group("/artists")
		{
			artists.GET(":artistID", h.GetArtistUser)
			artists.POST("", middleware.AuthRequired(), h.PostArtistUser)
			artists.PUT("", middleware.AuthRequired(), h.PutArtistUser)
			artists.DELETE("", middleware.AuthRequired(), h.DeleteArtistUser)
		}

		events := v1.Group("/events")
		{
			events.GET("/:eventID", h.GetEvent)
			events.GET("/:eventID/users", h.GetUserMetaByEventID)
			events.GET("/:eventID/invites", h.GetInvitesByEventID)
			events.POST("", middleware.AuthRequired(), h.PostEvent)
			events.PUT("", middleware.AuthRequired(), h.PutEvent)
			events.DELETE("", middleware.AuthRequired(), h.DeleteEvent)

		}

		invites := v1.Group("/invites")
		{
			invites.GET("", middleware.AuthRequired(), h.GetInvitesByArtistID)
			invites.GET("/:inviteID", middleware.AuthRequired(), h.GetInvite)
			invites.POST("", middleware.AuthRequired(), h.PostInvite)
			invites.PUT("", middleware.AuthRequired(), h.PutEvent)
			invites.PUT("/:inviteID/status", middleware.AuthRequired(), h.PutInviteStatus)
			invites.DELETE("", middleware.AuthRequired(), h.DeleteEvent)
		}

		tags := v1.Group("/tags")
		{
			tags.GET("", h.GetAllTags)
		}

		statics := v1.Group("/statics")
		{
			statics.POST("/images", middleware.AuthRequired(), h.PostImage)
		}

	}

	// OAuth用。 APIではないので注意
	auth := r.Group("/auth")
	{
		auth.GET("/:provider", h.AuthBegin)
		auth.GET("/:provider/callback", h.AuthCallback)
	}
	r.GET("/logout", middleware.AuthRequired(), h.AuthLogout)
}
