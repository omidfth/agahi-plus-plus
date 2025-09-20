package repository

import "agahi-plus-plus/internal/model"

type AppLogRepository interface {
	Insert(log *model.AppLog) error
}
