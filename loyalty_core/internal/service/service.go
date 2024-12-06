package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"loyalty_core/internal/service/business"
)

type Service struct {
	logger *zap.Logger
	gormDb *gorm.DB
	biz    *business.Business
}

func NewService(
	logger *zap.Logger,
	gormDb *gorm.DB,
) (*Service, error) {

	biz := business.NewBusiness(logger, gormDb)

	return &Service{
		logger: logger,
		gormDb: gormDb,
		biz:    biz,
	}, nil
}
