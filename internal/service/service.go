package service

import (
	"multi-language-event/config"
	store "multi-language-event/internal/storage"
)

type APIService struct {
	config *config.Config
}

func NewAPIService(s *store.Store, cfg *config.Config) *APIService {
	return &APIService{
		config: cfg,
	}
}
