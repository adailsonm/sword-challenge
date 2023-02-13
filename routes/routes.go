package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewTaskRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

// Route interface
type Route interface {
	Setup()
}

func NewRoutes(
	userRoutes UserRoutes,
	taskRoutes TaskRoutes,

) Routes {
	return Routes{
		userRoutes,
		taskRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
