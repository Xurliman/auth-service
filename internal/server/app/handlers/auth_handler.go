package handlers

import (
	"github.com/Xurliman/auth-service/internal/server/app/dto"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/pkg/json"
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

func (a AuthHandler) Login(ctx *fiber.Ctx) error {
	request := new(dto.LoginRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return json.Success(ctx, a.service.Login(request))
}

func (a AuthHandler) Register(ctx *fiber.Ctx) error {
	request := new(dto.RegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return json.Success(ctx, a.service.Register(request))
}
