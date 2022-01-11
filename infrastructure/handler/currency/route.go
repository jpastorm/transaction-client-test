package currency

import (
	"github.com/jpastorm/transaction-client-test/domain/currency"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/response"
	currencyStorage "github.com/jpastorm/transaction-client-test/infrastructure/postgres/currency"
	"github.com/jpastorm/transaction-client-test/model"

	"github.com/labstack/echo/v4"
)

// NewRouterCurrency returns a router to handle model.Currency requests
func NewCurrencyRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	// build middlewares to validate permissions on the routes

	adminRoutes(specification.Api, handler)
	privateRoutes(specification.Api, handler)
	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	responser := response.New(specification.Logger)

	useCase := currency.New(currencyStorage.New(specification.DB))
	return newHandler(useCase, responser)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("api/v1/admin/currency", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("/api/v1/private/currency", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(api *echo.Echo, h handler) {
	route := api.Group("/api/v1/public/currency")

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}
