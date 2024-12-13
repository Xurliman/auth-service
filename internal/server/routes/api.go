package routes

import (
	"github.com/Xurliman/auth-service/internal/database"
	"github.com/Xurliman/auth-service/internal/server/app/handlers"
	"github.com/Xurliman/auth-service/internal/server/app/middlewares"
	"github.com/Xurliman/auth-service/internal/server/app/repositories"
	"github.com/Xurliman/auth-service/internal/server/app/services"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	db := database.GetDB()
	apiRouter := app.Group("/api", middlewares.JwtMiddleware(db))

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	auth := apiRouter.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	auth.Post("/register", userHandler.Add)

	users := apiRouter.Group("/users")
	users.Get("/", userHandler.List)
	users.Get("/me", userHandler.GetMe)
	users.Get("/:id", userHandler.Show)
	users.Patch("/", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)
}
