package commands

import (
	"log"
	"os"

	"github.com/adailsonm/desafio-sword/lib"
	"github.com/adailsonm/desafio-sword/routes"
	"github.com/spf13/cobra"
)

type ServerCommand struct{}

func (s *ServerCommand) Short() string {
	return "server application"
}

func (s *ServerCommand) Setup(cmd *cobra.Command) {}

func (s *ServerCommand) Run() lib.CommandRunner {
	return func(
		router lib.RequestHandler,
		route routes.Routes,
		database lib.Database,
	) {
		route.Setup()

		log.Print("Running server")
		port := os.Getenv("PORT")
		if port == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + port)
		}
	}
}

func NewServerCommand() *ServerCommand {
	return &ServerCommand{}
}
