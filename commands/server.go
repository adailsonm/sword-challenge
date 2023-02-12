package commands

import (
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
		env lib.Env,
		router lib.RequestHandler,
		route routes.Routes,
		logger lib.Logger,
		database lib.Database,
	) {
		route.Setup()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}

func NewServerCommand() *ServerCommand {
	return &ServerCommand{}
}
