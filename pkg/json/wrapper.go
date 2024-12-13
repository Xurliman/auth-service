package json

import (
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/Xurliman/auth-service/pkg/pagination"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type DefaultResponse struct {
	Success  bool                   `json:"success"`
	Status   int                    `json:"status"`
	Code     string                 `json:"code,omitempty"`
	Message  string                 `json:"message"`
	Data     interface{}            `json:"data,omitempty"`
	Paginate *pagination.Pagination `json:"paginate,omitempty"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"description"`
	Message          string `json:"message"`
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(DefaultResponse{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func Pagination(ctx *fiber.Ctx, pagination *pagination.Pagination) error {
	return ctx.Status(fiber.StatusOK).JSON(DefaultResponse{
		Success:  true,
		Status:   fiber.StatusOK,
		Message:  "OK",
		Data:     pagination.Rows,
		Paginate: pagination,
	})
}

func Error(ctx *fiber.Ctx, err error, code string) error {
	errorMessage := logErrorFormat(err, code)
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusBadRequest).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusBadRequest,
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	})
}

func ErrorInternal(ctx *fiber.Ctx, err error, code string) error {
	errorMessage := logErrorFormat(err, code)
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusInternalServerError).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusInternalServerError,
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	})
}

func ErrorValidation(ctx *fiber.Ctx, err error) error {
	errorMessage := logErrorFormat(err, "E_VALIDATION")
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusBadRequest).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusBadRequest,
		Code:    "E_VALIDATION",
		Message: err.Error(),
		Data:    nil,
	})
}

func ErrorNotFound(ctx *fiber.Ctx, err error) error {
	errorMessage := logErrorFormat(err, "E_NOT_FOUND")
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusNotFound).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusNotFound,
		Code:    "E_NOT_FOUND",
		Message: err.Error(),
		Data:    nil,
	})
}

func ErrorUnauthorized(ctx *fiber.Ctx, err error) error {
	errorMessage := logErrorFormat(err, "E_UNAUTHORIZED")
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusUnauthorized).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusUnauthorized,
		Code:    "E_UNAUTHORIZED",
		Message: err.Error(),
		Data:    nil,
	})
}

func ErrorForbidden(ctx *fiber.Ctx, err error) error {
	errorMessage := logErrorFormat(err, "E_FORBIDDEN")
	log.Error(errorMessage, zap.Error(err))
	return ctx.Status(fiber.StatusForbidden).JSON(DefaultResponse{
		Success: false,
		Status:  fiber.StatusForbidden,
		Code:    "E_FORBIDDEN",
		Message: err.Error(),
		Data:    nil,
	})
}

func logErrorFormat(err error, code string) string {
	return "❌ " + "[" + code + "] " + err.Error()
}
