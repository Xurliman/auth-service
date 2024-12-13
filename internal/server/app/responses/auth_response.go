package responses

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"time"
)

type AuthResponse struct {
	Id              string    `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       string    `json:"created_at"`
	ExpiresAt       time.Time `json:"expires_at"`
}

func AuthLoginTransformer(user models.User, expiresAt int64) AuthResponse {
	return AuthResponse{
		Id:              user.Id.String(),
		Username:        user.Username,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		CreatedAt:       user.CreatedAt.Format(constants.TimestampFormat),
		ExpiresAt:       time.Unix(expiresAt, 0),
	}
}
