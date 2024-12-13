package repositories

import (
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) Insert(user models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func NewAuthRepository(db *gorm.DB) interfaces.IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}
