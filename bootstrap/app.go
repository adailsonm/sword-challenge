package bootstrap

import (
	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/internal/repository"
	"github.com/adailsonm/desafio-sword/internal/services"
	"github.com/adailsonm/desafio-sword/lib"
	"github.com/adailsonm/desafio-sword/middlewares"
	"github.com/adailsonm/desafio-sword/routes"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	middlewares.Module,
	routes.Module,
	lib.Module,
	services.Module,
	repository.Module,
)
