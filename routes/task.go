package routes

import (
	"log"

	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/lib"
	"github.com/adailsonm/desafio-sword/middlewares"
)

type TaskRoutes struct {
	handler        lib.RequestHandler
	taskController controllers.TaskController
	authMiddleware middlewares.AuthMiddleware
}

func (s TaskRoutes) Setup() {
	log.Print("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/tasks", s.taskController.GetTask)
		api.GET("/tasks/:id", s.taskController.GetOneTask)
		api.POST("/tasks", s.taskController.SaveTask)
		api.POST("/tasks/:id", s.taskController.UpdateTask)
		api.DELETE("/tasks/:id", s.taskController.DeleteTask)
	}
}

func NewTaskRoutes(
	handler lib.RequestHandler,
	taskController controllers.TaskController,
	authMiddleware middlewares.AuthMiddleware,

) TaskRoutes {
	return TaskRoutes{
		handler:        handler,
		taskController: taskController,
		authMiddleware: authMiddleware,
	}
}
