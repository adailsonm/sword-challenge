package bootstrap

import (
	"github.com/adailsonm/desafio-sword/internal/controllers"
	"github.com/adailsonm/desafio-sword/routes"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
)
