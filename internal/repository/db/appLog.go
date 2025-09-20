package db

import (
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewAppLog(db *gorm.DB, logger *zap.Logger) repository.AppLogRepository {
	return &appLog{
		db:     db,
		logger: logger,
	}
}

type appLog struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (a appLog) Insert(log *model.AppLog) error {
	return a.db.Create(log).Error
}
