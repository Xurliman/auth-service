package handlers

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/Xurliman/auth-service/pkg/json"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service interfaces.IAuthService
}

func NewAuthHandler(service interfaces.IAuthService) interfaces.IAuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	log.Info("✅ AUTH LOGIN")
	request := new(requests.LoginRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return json.Error(ctx, err, "ERR_INVALID_REQUEST")
	}

	return h.service.Login(ctx, request)
}

func (h *AuthHandler) Logout(ctx *fiber.Ctx) error {
	log.Info("✅ AUTH LOGOUT")
	return h.service.Logout(ctx)
}

func (h *AuthHandler) VerifyEmail(ctx *fiber.Ctx) error {
	tokenString := ctx.Query("token")
	if tokenString == "" {
		return json.Error(ctx, constants.ErrInvalidToken, "ERR_INVALID_REQUEST")
	}

	return h.service.VerifyEmail(ctx, tokenString)
}
