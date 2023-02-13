package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabaseTrx),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	dbTrxMiddleware DatabaseTrx,
	authMiddleware AuthMiddleware,
) Middlewares {
	return Middlewares{
		dbTrxMiddleware,
		authMiddleware,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
