package interfaces

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/Xurliman/auth-service/pkg/pagination"
	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	Add(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	GetMe(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type IUserService interface {
	Add(ctx *fiber.Ctx, request *requests.StoreUserRequest) error
	List(ctx *fiber.Ctx, paginate pagination.Pagination) error
	Show(ctx *fiber.Ctx, id string) error
	GetMe(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx, id string, request *requests.UpdateUserRequest) error
	Delete(ctx *fiber.Ctx, id string) error
}

type IUserRepository interface {
	Create(user models.User) (models.User, error)
	EmailExists(email string) bool
	FindByEmail(email string) (user models.User, err error)
	FindById(id string) (user models.User, err error)
	GetAll(pagination pagination.Pagination) (*pagination.Pagination, error)
	UpdateById(id string, storeData models.User) (models.User, error)
	Delete(id string) error
}
