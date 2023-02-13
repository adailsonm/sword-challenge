package middlewares

import (
	"net/http"
	"strings"

	"github.com/adailsonm/desafio-sword/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service services.AuthService
}

func NewAuthMiddleware(
	service services.AuthService,
) AuthMiddleware {
	return AuthMiddleware{
		service: service,
	}
}

func (a AuthMiddleware) Setup() {}

func (a AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authorizationHeader, " ")
		if len(t) == 2 {
			token := t[1]
			isAuthorized, err := a.service.Authorize(token)
			if isAuthorized {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}
}
