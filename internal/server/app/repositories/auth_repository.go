package repositories

import (
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) FindByEmail(email string) (user models.User, err error) {
	if err = r.db.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) MakeEmailVerified(email string) (err error) {
	if err = r.db.
		Model(&models.User{}).
		Where("email = ?", email).
		Update("is_email_verified", true).
		Error; err != nil {
		return err
	}
	
	return nil
}

func (r *AuthRepository) AddSession(session models.UserSession) (models.UserSession, error) {
	if err := r.db.
		Model(&models.UserSession{}).
		Where("is_active = ? AND user_id = ?", true, session.UserId).
		Update("is_active", false).
		Error; err != nil {
		return session, err
	}

	if err := r.db.
		Create(&session).
		Error; err != nil {
		return session, err
	}
	return session, nil
}

func (r *AuthRepository) FindSessionByToken(token string) (models.UserSession, error) {
	var session models.UserSession
	if err := r.db.
		Where("session_token = ?", token).
		First(&session).
		Error; err != nil {
		return session, err
	}
	return session, nil
}

func (r *AuthRepository) UpdateSession(id string, session models.UserSession) (models.UserSession, error) {
	var userSession models.UserSession
	if err := r.db.
		Model(&userSession).
		Where("id = ?", id).
		Updates(&session).
		Error; err != nil {
		return userSession, err
	}
	return userSession, nil
}

func (r *AuthRepository) MakeSessionInactive(id string) error {
	var userSession models.UserSession
	if err := r.db.
		Model(&userSession).
		Where("id = ?", id).
		Update("is_active", false).
		Error; err != nil {
		return err
	}
	return nil
}
