package routes

import (
	"github.com/Xurliman/auth-service/internal/database"
	"github.com/Xurliman/auth-service/internal/server/app/handlers"
	"github.com/Xurliman/auth-service/internal/server/app/repositories"
	"github.com/Xurliman/auth-service/internal/server/app/services"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db := database.GetDB()

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	api := app.Group("/api")
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api.Post("/logout", authHandler.Login)
	api.Get("/me", authHandler.Login)
}
