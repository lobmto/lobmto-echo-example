package server

import (
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, app *App) {
	healthController := app.healthController
	e.GET("/health", healthController.Check)

	wordController := app.wordController
	e.GET("/words/:id", wordController.GetWord)
	e.POST("/words", wordController.PostWord)
}
