package interfaces

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/gofiber/fiber/v2"
)

type IAuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx) error
}

type IAuthService interface {
	Login(ctx *fiber.Ctx, request *requests.LoginRequest) error
	Logout(ctx *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx, tokenString string) error
}

type IAuthRepository interface {
	FindByEmail(email string) (user models.User, err error)
	MakeEmailVerified(email string) (err error)
	AddSession(session models.UserSession) (models.UserSession, error)
	FindSessionByToken(token string) (models.UserSession, error)
	UpdateSession(id string, session models.UserSession) (models.UserSession, error)
	MakeSessionInactive(id string) error
}
