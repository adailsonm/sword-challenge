package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabaseTrx),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
	dbTrxMiddleware DatabaseTrx,
) Middlewares {
	return Middlewares{
		dbTrxMiddleware,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
