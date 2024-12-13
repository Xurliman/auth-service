package interfaces

import (
	"github.com/Xurliman/auth-service/internal/server/app/dto"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/gofiber/fiber/v2"
)

type IAuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type IAuthService interface {
	Register(request *dto.RegisterRequest) error
	Login(request *dto.LoginRequest) error
}

type IAuthRepository interface {
	Insert(user models.User) error
}
