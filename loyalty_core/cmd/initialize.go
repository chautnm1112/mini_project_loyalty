package main

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logr "gorm.io/gorm/logger"
	oslog "log"
	"loyalty_core/config"
	"loyalty_core/internal/service"
	"os"
	"time"
)

const (
	DefaultMaxIdlesConst = 10
	DefaultMaxOpenConst
	DefaultSlowThreshold = 10 * time.Second
)

func newService(cfg *config.Config) (*service.Service, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=mode\n", cfg.Username, cfg.Password, cfg.DbHost, cfg.Port, cfg.DBName)

	gormDb, err := ConnectPostgresql(dsn)
	if err != nil {
		logger.Error("Error init gormDb", zap.Error(err))
		return nil, err
	}

	return service.NewService(logger, gormDb)
}

func ConnectPostgresql(dsn string) (*gorm.DB, error) {
	newLogger := gorm_logr.New(
		oslog.New(os.Stderr, "", oslog.LstdFlags),
		gorm_logr.Config{
			SlowThreshold: DefaultSlowThreshold,
			LogLevel:      gorm_logr.Warn,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(DefaultMaxIdlesConst)
	sqlDB.SetMaxOpenConns(DefaultMaxOpenConst)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		return nil, err
	}

	return db, nil
}
