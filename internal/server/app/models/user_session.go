package models

import (
	"github.com/google/uuid"
	"time"
)

type UserSession struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId       uuid.UUID `json:"user_id" gorm:"not null"`
	SessionToken string    `json:"session_token" gorm:"not null"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`

	User User `json:"user"`
}

func (m *UserSession) TableName() string {
	return "user_sessions"
}
