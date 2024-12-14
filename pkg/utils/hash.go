package utils

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), constants.HashingCost)
	if err != nil {
		log.Warn("error hashing password", zap.Error(err))
		return ""
	}
	return string(bytes)
}
