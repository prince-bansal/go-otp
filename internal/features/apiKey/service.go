package apiKey

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/features/apiKey/domain"
)

type ApiService interface {
	GetAll(ctx context.Context, id int) ([]*domain.ApiKeyD, error)
	Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiKeyD, error)
	Expire(ctx context.Context, id int) (*domain.ApiKeyD, error)
}

func NewApiService(repository ApiKeyRepository) ApiService {
	return &ApiServiceImpl{
		repository: repository,
	}
}

type ApiServiceImpl struct {
	repository ApiKeyRepository
}

func (s *ApiServiceImpl) GetAll(ctx context.Context, id int) ([]*domain.ApiKeyD, error) {
	records, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *ApiServiceImpl) Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiKeyD, error) {
	createdRecord, err := s.repository.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return createdRecord, nil
}

func (s *ApiServiceImpl) Expire(ctx context.Context, id int) (*domain.ApiKeyD, error) {
	record, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	deletedRecord, err := s.repository.SoftDelete(ctx, record)
	return deletedRecord, err
}
