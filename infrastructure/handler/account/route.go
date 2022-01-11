package account

import (
	"github.com/jpastorm/transaction-client-test/domain/account"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/response"
	accountStorage "github.com/jpastorm/transaction-client-test/infrastructure/postgres/account"
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/labstack/echo/v4"
)

// NewAccountRouter returns a router to handle model.Account requests
func NewAccountRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	// build middlewares to validate permissions on the routes

	adminRoutes(specification.Api, handler)
	privateRoutes(specification.Api, handler)
	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	responser := response.New(specification.Logger)

	useCase := account.New(accountStorage.New(specification.DB))
	return newHandler(useCase, responser)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("api/v1/admin/account", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("/api/v1/private/account", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(api *echo.Echo, h handler) {
	route := api.Group("/api/v1/public/account")
	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}
