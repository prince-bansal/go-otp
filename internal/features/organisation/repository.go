package organisation

import (
	"context"
	"fmt"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"github.com/prince-bansal/go-otp/models"
	"time"
)

type OrganisationRepository interface {
	GetAll(ctx context.Context) (*[]domain.OrganisationD, error)
	GetOne(ctx context.Context, id string) (*domain.OrganisationD, error)
	Register(ctx context.Context, request domain.OrganisationRequest) (*domain.OrganisationD, error)
}

type OrganisationRepositoryImpl struct {
}

func NewOrganisationRepository() OrganisationRepository {
	return &OrganisationRepositoryImpl{}
}

func (r *OrganisationRepositoryImpl) GetAll(ctx context.Context) (*[]domain.OrganisationD, error) {
	return &Organisations, nil
}

func (r *OrganisationRepositoryImpl) GetOne(ctx context.Context, id string) (*domain.OrganisationD, error) {
	println("organisation>>>>", len(Organisations))
	return &Organisations[0], nil
}

func (r *OrganisationRepositoryImpl) Register(ctx context.Context, request domain.OrganisationRequest) (*domain.OrganisationD, error) {
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
	return &newOrg, nil
}
