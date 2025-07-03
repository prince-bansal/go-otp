package api_key

import (
	"context"
	domain "github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/pkg/logger"
)

type ApiService interface {
	GetAll(ctx context.Context, id int) ([]*domain.ApiKeyD, error)
	Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiGenerateResponse, error)
	Expire(ctx context.Context, id int) (*domain.ApiKeyD, error)
	GetByApiKey(ctx context.Context, api string) (*domain.ApiKeyD, error)
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
		logger.Error("failed to get the api key(s)", err)
		return nil, err
	}
	return records, nil
}

func (s *ApiServiceImpl) Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiGenerateResponse, error) {

	request.GenerateKey()
	response := &domain.ApiGenerateResponse{
		Otp: request.Key,
	}

	err := request.HashKey()
	if err != nil {
		logger.Error("failed to hash key: %s", request.Key)
		return nil, err
	}

	_, err = s.repository.Create(ctx, request)
	if err != nil {
		logger.Error("failed to save user", err)
		return nil, err
	}
	return response, nil
}

func (s *ApiServiceImpl) Expire(ctx context.Context, id int) (*domain.ApiKeyD, error) {
	record, err := s.repository.FindById(ctx, id)
	if err != nil {
		logger.Error("api key not found: %d", id)
		return nil, err
	}
	deletedRecord, err := s.repository.SoftDelete(ctx, record)
	return deletedRecord, err
}

func (s *ApiServiceImpl) GetByApiKey(ctx context.Context, apiKey string) (*domain.ApiKeyD, error) {

	d := &domain.ApiKeyD{Key: apiKey}
	err := d.HashKey()
	if err != nil {
		logger.Error("failed to hash the key: %s", d.Key)
		return nil, err
	}

	record, err := s.repository.GetBySaltHash(ctx, d.Salt)
	if err != nil {
		logger.Error("failed to generate salt for key: %s", apiKey)
		return nil, err
	}

	success, err := d.CompareKey(apiKey)

	if err != nil || !success {
		logger.Error("key mismatch", err)
		return nil, err
	}

	return record, nil
}
