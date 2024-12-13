package database

import (
	"fmt"
	"github.com/Xurliman/auth-service/internal/config/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func Setup(cfg *config.DatabaseSettings) error {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Name,
		cfg.Password,
		cfg.Port,
		cfg.SSLMode,
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	gormDB = db

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to retrieve sql.db from gorm: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.MaxConnIdleTime)

	for retries := 0; retries < 3; retries++ {
		if err = sqlDB.Ping(); err == nil {
			log.Println("Successfully connected to the database")
			break
		} else if retries == 2 {
			return fmt.Errorf("failed to connect to the database after retries: %w", err)
		} else {
			log.Printf("Retrying database connection... (%d/3)\n", retries+1)
			time.Sleep(2 * time.Second)
		}
	}
	return nil
}

func CloseDB() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to retrieve sql.db instance: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
