package organisation

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
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
		return nil, err
	}
	return records, nil
}

func (s *OrganisationServiceImpl) GetOne(ctx context.Context, id string) (*domain.OrganisationD, error) {
	record, err := s.organisationRepository.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *OrganisationServiceImpl) Register(ctx context.Context, request *domain.OrganisationD) (*domain.OrganisationD, error) {
	createdRecord, err := s.organisationRepository.Register(ctx, request)
	if err != nil {
		return nil, err
	}
	return createdRecord, nil
}

func (s *OrganisationServiceImpl) GetByApiKey(ctx context.Context, api string) (*domain.OrganisationD, error) {
	record, err := s.organisationRepository.GetByApiKey(ctx, api)
	if err != nil {
		return nil, err
	}
	return record, nil
}
