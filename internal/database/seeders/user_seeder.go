package seeders

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/repositories"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
}

func (s UserSeeder) Seed(db *gorm.DB) error {
	users := []models.User{
		{
			Name:     "Admin",
			Username: "admin",
			Email:    "admin@example.com",
			Password: hashPassword("123"),
		},
	}
	repo := repositories.NewUserRepository(db)
	for _, user := range users {
		_, err := repo.Create(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Warn("error hashing password", zap.Error(err))
		return ""
	}
	return string(bytes)
}
