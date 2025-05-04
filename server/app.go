package server

import (
	"lobmto-echo-example/controllers"
	"lobmto-echo-example/infrastructure/repositories"

	"gorm.io/gorm"
)

type App struct {
	healthController *controllers.HealthController
	wordController   *controllers.WordController
}

func NewApp(db *gorm.DB) *App {
	wordRepository := repositories.NewWordRepository(db)
	return &App{
		healthController: controllers.NewHealthController(),
		wordController:   controllers.NewWordController(wordRepository),
	}
}
