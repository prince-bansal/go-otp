package organisation

import (
	"context"
	"fmt"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"github.com/prince-bansal/go-otp/models"
	"time"
)

type OrganisationService interface {
	GetAll(ctx context.Context) []domain.OrganisationD
	GetOne(ctx context.Context, id string) domain.OrganisationD
	Register(ctx context.Context, request domain.OrganisationRequest) domain.OrganisationD
}

func NewOrganisationService() OrganisationService {
	return &OrganisationServiceImpl{
		records: Organisations,
	}
}

type OrganisationServiceImpl struct {
	records []domain.OrganisationD
}

func (s *OrganisationServiceImpl) GetAll(ctx context.Context) []domain.OrganisationD {
	return Organisations
}

func (s *OrganisationServiceImpl) GetOne(ctx context.Context, id string) domain.OrganisationD {
	// dummy implementation
	// todo: implement real method
	return Organisations[0]
}

func (s *OrganisationServiceImpl) Register(ctx context.Context, request domain.OrganisationRequest) domain.OrganisationD {
	id := len(Organisations) + 1
	newOrg := domain.OrganisationD{
		Id:        fmt.Sprintf("%d", id),
		Name:      fmt.Sprintf("Organisation %d", id),
		Email:     fmt.Sprintf("org%d@email.com", id),
		Package:   models.SILVER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	Organisations = append(Organisations, newOrg)
	return newOrg
}
