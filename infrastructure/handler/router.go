package handler

import (
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/account"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/client"
	"github.com/jpastorm/transaction-client-test/infrastructure/handler/currency"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/jpastorm/transaction-client-test/model"
)

// InitRoutes initialize all the routes of the service
func InitRoutes(specification model.RouterSpecification) {
	// initialize auth package to validate tokens
	// initialize scope package to validate permissions on admin routes

	// A
	account.NewAccountRouter(specification)
	// B
	// C
	client.NewClientRouter(specification)
	currency.NewCurrencyRouter(specification)
	// D
	// E
	// F
	// G
	// H
	healthRoute(specification.Api)
	// I
	// J
	// K
	// L
	// M
	// N
	// O
	// P
	// Q
	// R
	// S
	// T
	// U
	// V
	// W
	// X
	// Y
	// Z
}

func healthRoute(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"service_name": "",
			},
		)
	})
}
