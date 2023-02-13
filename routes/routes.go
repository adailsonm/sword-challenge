package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewTaskRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	userRoutes UserRoutes,
	taskRoutes TaskRoutes,
	authRoutes AuthRoutes,

) Routes {
	return Routes{
		userRoutes,
		taskRoutes,
		authRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
