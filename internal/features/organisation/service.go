package organisation

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/pkg/logger"
)

type OrganisationService interface {
	GetAll(ctx context.Context) ([]*domain.OrganisationD, error)
	GetOne(ctx context.Context, id string) (*domain.OrganisationD, error)
	Register(ctx context.Context, request *domain.OrganisationD) (*domain.OrganisationD, error)
	GetByApiKey(ctx context.Context, api string) (*domain.OrganisationD, error)
}

func NewOrganisationService(organisationRepository OrganisationRepository) OrganisationService {
	return &OrganisationServiceImpl{
		organisationRepository: organisationRepository,
	}
}

type OrganisationServiceImpl struct {
	organisationRepository OrganisationRepository
}

func (s *OrganisationServiceImpl) GetAll(ctx context.Context) ([]*domain.OrganisationD, error) {
	records, err := s.organisationRepository.GetAll(ctx)
	if err != nil {
		logger.Error("failed to fetch organisations", err)
		return nil, err
	}
	return records, nil
}

func (s *OrganisationServiceImpl) GetOne(ctx context.Context, id string) (*domain.OrganisationD, error) {
	record, err := s.organisationRepository.GetOne(ctx, id)
	if err != nil {
		logger.Error("organisation not found for id %s", id, err)
		return nil, err
	}
	return record, nil
}

func (s *OrganisationServiceImpl) Register(ctx context.Context, request *domain.OrganisationD) (*domain.OrganisationD, error) {
	createdRecord, err := s.organisationRepository.Register(ctx, request)
	if err != nil {
		logger.Error("failed to save organisation", err)
		return nil, err
	}
	return createdRecord, nil
}

func (s *OrganisationServiceImpl) GetByApiKey(ctx context.Context, api string) (*domain.OrganisationD, error) {
	record, err := s.organisationRepository.GetByApiKey(ctx, api)
	if err != nil {
		logger.Error("organisation not found with api key %s", api, err)
		return nil, err
	}
	return record, nil
}
