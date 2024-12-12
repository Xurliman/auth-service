package main

import (
	"github.com/Xurliman/auth-service/internal/config/config"
	"log"
)

func main() {
	cfg := config.Setup()
	log.Print(cfg.Database.Name)
}
