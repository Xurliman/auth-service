package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email     string         `gorm:"type:varchar(255);unique_index;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time      `gorm:"DEFAULT:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"DEFAULT:null"`
}