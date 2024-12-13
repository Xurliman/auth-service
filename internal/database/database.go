package database

import (
	"fmt"
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"time"
)

var (
	db *gorm.DB
)

func Setup(cfg *config.AppConfig) error {
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.User,
		dbCfg.Name,
		dbCfg.Password,
		dbCfg.Port,
		dbCfg.SSLMode,
	)

	gormLogger := zapgorm2.New(log.GetSQLLogger(cfg.App.Env))
	gormLogger.SetAsDefault()
	gormLogger.SlowThreshold = 200 * time.Millisecond
	gormLogger.IgnoreRecordNotFoundError = true
	if cfg.App.Debug {
		gormLogger.LogLevel = logger.Info
	} else {
		gormLogger.LogLevel = logger.Warn
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	db = gormDB

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to retrieve sql.db from gorm: %w", err)
	}

	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(dbCfg.MaxConnLifetime)
	sqlDB.SetConnMaxIdleTime(dbCfg.MaxConnIdleTime)

	for retries := 0; retries < 3; retries++ {
		if err = sqlDB.Ping(); err == nil {
			log.Info("Successfully connected to the database")
			break
		} else if retries == 2 {
			return fmt.Errorf("failed to connect to the database after retries: %w", err)
		} else {
			log.Warn("Retrying database connection... (%d/3)\n", zap.Int("retries", retries+1))
			time.Sleep(2 * time.Second)
		}
	}
	return nil
}

func CloseDB() error {
	log.Info("closing database connection...")
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
