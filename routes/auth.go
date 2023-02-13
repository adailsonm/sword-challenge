package routes

import (
	"log"

	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/lib"
)

type AuthRoutes struct {
	handler        lib.RequestHandler
	authController controllers.AuthController
}

func (s AuthRoutes) Setup() {
	log.Print("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.POST("/auth/login", s.authController.Login)
	}
}

func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		authController: authController,
	}
}
