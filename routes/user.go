package routes

import (
	"log"

	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/lib"
	"github.com/adailsonm/desafio-sword/middlewares"
)

type UserRoutes struct {
	handler        lib.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.AuthMiddleware
}

func (s UserRoutes) Setup() {
	log.Print("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/users", s.userController.GetUser)
		api.GET("/users/:id", s.userController.GetOneUser)
		api.POST("/users", s.userController.SaveUser)
		api.POST("/users/:id", s.userController.UpdateUser)
		api.DELETE("/users/:id", s.userController.DeleteUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	handler lib.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.AuthMiddleware,

) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
