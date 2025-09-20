package repository

import (
	"agahi-plus-plus/internal/model"
)

type PostApiRepo interface {
	Get(token string, ServiceName string) (*model.Post, error)
}

type PostDbRepo interface {
	Insert(*model.Post) (*model.Post, error)
	Get(token string) (*model.Post, error)
	Update(post *model.Post) error
}
