package server

import (
	"lobmto-echo-example/controllers"
	"lobmto-echo-example/infrastructure/repositories"

	"gorm.io/gorm"
)

type App struct {
	healthController *controllers.HealthController
	wordsController   *controllers.WordsController
}

func NewApp(db *gorm.DB) *App {
	wordsRepository := repositories.NewWordsRepository(db)
	return &App{
		healthController: controllers.NewHealthController(),
		wordsController:   controllers.NewWordController(wordsRepository),
	}
}
