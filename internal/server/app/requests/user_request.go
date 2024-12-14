package requests

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/pkg/utils"
	"github.com/Xurliman/auth-service/pkg/validate"
)

type StoreUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *StoreUserRequest) Validate() error {
	return validate.ExtractValidationError(r)
}

func (r *StoreUserRequest) ToModel() models.User {
	return models.User{
		Name:     r.Name,
		Username: r.Username,
		Email:    r.Email,
		Password: utils.HashPassword(r.Password),
	}
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *UpdateUserRequest) Validate() error {
	return validate.ExtractValidationError(r)
}

func (r *UpdateUserRequest) ToModel() models.User {
	return models.User{
		Name:     r.Name,
		Username: r.Username,
		Email:    r.Email,
		Password: utils.HashPassword(r.Password),
	}
}
