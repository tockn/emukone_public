package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("id")
		if v == nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		_, ok := v.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
	}
}
