package repository

import "agahi-plus-plus/internal/model"

type ConfigRepository interface {
	List(serviceName string) []model.Config
	GetByCodes(codes []string, serviceName string) []model.Config
	ListAsMap(serviceName string) map[string]string
}
