package seeders

import (
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/repositories"
	"github.com/Xurliman/auth-service/pkg/utils"
	"gorm.io/gorm"
)

type UserSeeder struct {
}

func (s UserSeeder) Seed(db *gorm.DB) error {
	users := []models.User{
		{
			Name:     "Khurliman",
			Username: "admin",
			Email:    "jumamuratovahurliman8@gmail.com",
			Password: utils.HashPassword("123"),
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
