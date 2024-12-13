package handlers

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/Xurliman/auth-service/pkg/json"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/Xurliman/auth-service/pkg/pagination"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service interfaces.IUserService
}

func NewUserHandler(service interfaces.IUserService) interfaces.IUserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Add(ctx *fiber.Ctx) error {
	log.Info("✅ USER CREATE")

	user := new(requests.StoreUserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return json.Error(ctx, err, "ERR_INVALID_REQUEST")
	}

	return h.service.Add(ctx, user)
}

func (h *UserHandler) List(ctx *fiber.Ctx) error {
	log.Info("✅ USER LIST")

	paginate := pagination.GetPaginationParams(ctx)
	return h.service.List(ctx, paginate)
}

func (h *UserHandler) Show(ctx *fiber.Ctx) error {
	log.Info("✅ USER SHOW")

	id := ctx.Params("id")
	if id == "" {
		return json.ErrorValidation(ctx, constants.ErrIdRequired)
	}

	return h.service.Show(ctx, id)
}

func (h *UserHandler) GetMe(ctx *fiber.Ctx) error {
	log.Info("✅ USER GET ME")

	return h.service.GetMe(ctx)
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	log.Info("✅ USER UPDATE")

	id := ctx.Params("id")
	if id == "" {
		return json.ErrorValidation(ctx, constants.ErrIdRequired)
	}

	user := new(requests.UpdateUserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return json.Error(ctx, err, "ERR_INVALID_REQUEST")
	}

	return h.service.Update(ctx, id, user)
}

func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	log.Info("✅ USER DELETE")

	id := ctx.Params("id")
	if id == "" {
		return json.ErrorValidation(ctx, constants.ErrIdRequired)
	}

	return h.service.Delete(ctx, id)
}
