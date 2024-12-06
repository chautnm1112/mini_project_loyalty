package business

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"loyalty_core/internal/repository"
)

type Business struct {
	db         *gorm.DB
	repository *repository.Repository
}

func NewBusiness(
	logger *zap.Logger,
	db *gorm.DB,
) *Business {
	repo := repository.NewRepository(db)
	return &Business{
		db:         db,
		repository: repo,
	}
}
