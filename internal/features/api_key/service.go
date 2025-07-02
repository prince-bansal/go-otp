package api_key

import (
	"context"
	domain "github.com/prince-bansal/go-otp/internal/domain"
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
		return nil, err
	}

	_, err = s.repository.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ApiServiceImpl) Expire(ctx context.Context, id int) (*domain.ApiKeyD, error) {
	record, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	deletedRecord, err := s.repository.SoftDelete(ctx, record)
	return deletedRecord, err
}

func (s *ApiServiceImpl) GetByApiKey(ctx context.Context, apiKey string) (*domain.ApiKeyD, error) {

	d := &domain.ApiKeyD{Key: apiKey}
	err := d.HashKey()
	if err != nil {
		return nil, err
	}

	record, err := s.repository.GetBySaltHash(ctx, d.Salt)
	if err != nil {
		return nil, err
	}

	success, err := d.CompareKey(apiKey)

	if err != nil || !success {
		return nil, err
	}

	return record, nil
}
