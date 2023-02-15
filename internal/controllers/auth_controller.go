package controllers

import (
	"net/http"

	"github.com/adailsonm/desafio-sword/internal/models"
	"github.com/adailsonm/desafio-sword/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
	}
}

func (a AuthController) Login(c *gin.Context) {
	credentials := models.Credentials{}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := a.userService.GetOneByEmail(credentials.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	accessToken, err := a.authService.CreateToken(&user, c)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})

}
