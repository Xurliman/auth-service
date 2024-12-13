package requests

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/Xurliman/auth-service/pkg/validate"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
		Password: hashPassword(r.Password),
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
		Password: hashPassword(r.Password),
	}
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Warn("error hashing password", zap.Error(err))
		return ""
	}
	return string(bytes)
}
