package main

import (
	"fmt"
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/routes"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Setup()
	log.InitLogger(cfg.App.Env, constants.LogPath)

	app := fiber.New(fiber.Config{})
	routes.Setup(app)

	err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatal("error starting server", zap.Error(err))
	}
}
