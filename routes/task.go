package routes

import (
	"log"

	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/lib"
)

type TaskRoutes struct {
	handler        lib.RequestHandler
	taskController controllers.TaskController
}

func (s TaskRoutes) Setup() {
	log.Print("Setting up routes")
	api := s.handler.Gin.Group("/api")
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
) TaskRoutes {
	return TaskRoutes{
		handler:        handler,
		taskController: taskController,
	}
}
