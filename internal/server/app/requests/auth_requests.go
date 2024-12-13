package requests

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/pkg/validate"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) Validate() error {
	return validate.ExtractValidationError(r)
}

func (r *LoginRequest) ToModel() models.User {
	return models.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
