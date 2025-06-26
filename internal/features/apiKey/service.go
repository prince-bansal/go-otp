package apiKey

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/features/apiKey/domain"
)

type ApiService interface {
	GetAll(ctx context.Context, id string) []domain.ApiKeyD
	Create(ctx context.Context, request domain.ApiKeyRequest) domain.ApiKeyD
}

func NewApiService() ApiService {
	return &ApiServiceImpl{}
}

type ApiServiceImpl struct {
}

func (s *ApiServiceImpl) GetAll(ctx context.Context, id string) []domain.ApiKeyD {
	return apiKeys
}

func (s *ApiServiceImpl) Create(ctx context.Context, request domain.ApiKeyRequest) domain.ApiKeyD {
	return apiKeys[0]
}
