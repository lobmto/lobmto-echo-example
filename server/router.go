package server

import (
	"lobmto-echo-example/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	healthController := controllers.NewHealthController()
	e.GET("/health", healthController.Check)

	return e
}
