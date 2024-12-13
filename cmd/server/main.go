package main

import (
	"fmt"
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/database"
	"github.com/Xurliman/auth-service/internal/database/seeders"
	"github.com/Xurliman/auth-service/internal/server/routes"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
)

func main() {
	cfg := config.Setup()
	log.InitLogger(cfg.App.Env, constants.LogPath)

	app := fiber.New(fiber.Config{})
	err := database.Setup(cfg)
	if err != nil {
		log.Fatal("error connecting to database", zap.Error(err))
	}
	defer func() {
		err = database.CloseDB()
		if err != nil {
			log.Fatal("error closing database connection", zap.Error(err))
		}
	}()

	argsListener()
	routes.Setup(app)
	err = app.Listen(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatal("error starting server", zap.Error(err))
	}
}

func argsListener() {
	db := database.GetDB()
	for _, arg := range os.Args {
		if arg == "--seed" {
			runner := seeders.All(db)
			if err := runner.Run(); err != nil {
				log.Fatal("error seeding database", zap.Error(err))
			}
		}
	}
}
