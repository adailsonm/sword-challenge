package routes

import (
	"log"

	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/lib"
)

type UserRoutes struct {
	handler        lib.RequestHandler
	userController controllers.UserController
}

func (s UserRoutes) Setup() {
	log.Print("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetUser)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.POST("/user", s.userController.SaveUser)
		api.POST("/user/:id", s.userController.UpdateUser)
		api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	handler lib.RequestHandler,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
