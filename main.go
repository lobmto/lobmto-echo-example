package main

import (
	"lobmto-echo-example/infrastructure/database"
	"lobmto-echo-example/infrastructure/migration"
	"lobmto-echo-example/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.InitDB()
	if err != nil {
		e.Logger.Fatalf("Failed to initialize database: %v", err)
	}

	// TODO: DI Container に移行したい
	app := server.NewApp(db)
	server.SetupRouter(e, app)

	// TODO: マイグレーション用のコマンドとして切り出す
	if err := migration.Migrate(db); err != nil {
		e.Logger.Fatalf("Failed to migrate database: %v", err)
	}

	e.Logger.Fatal(e.Start(":8000"))
}
