package main

import (
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Setup()
	log.InitLogger(cfg.App.Env, constants.LogPath)
	log.Info("Starting Auth Service", zap.String("some", "data"))
}
