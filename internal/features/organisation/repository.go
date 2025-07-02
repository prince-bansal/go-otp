package organisation

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/models"
	"github.com/prince-bansal/go-otp/store/db"
	"gorm.io/gorm"
)

type OrganisationRepository interface {
	GetAll(ctx context.Context) ([]*domain.OrganisationD, error)
	GetOne(ctx context.Context, id string) (*domain.OrganisationD, error)
	GetByApiKey(ctx context.Context, key string) (*domain.OrganisationD, error)
	Register(ctx context.Context, request *domain.OrganisationD) (*domain.OrganisationD, error)
}

type OrganisationRepositoryImpl struct {
	db *gorm.DB
}

func NewOrganisationRepository(store *db.Store) OrganisationRepository {
	return &OrganisationRepositoryImpl{
		db: store.Db,
	}
}

func (r *OrganisationRepositoryImpl) GetAll(ctx context.Context) ([]*domain.OrganisationD, error) {
	var records []models.Organisation
	if err := r.
		db.
		WithContext(ctx).
		Find(&records).Error; err != nil {
		return nil, err
	}

	var processedRecords []*domain.OrganisationD
	for _, m := range records {
		processedRecords = append(processedRecords, m.ToDomain())
	}
	return processedRecords, nil
}

func (r *OrganisationRepositoryImpl) GetOne(ctx context.Context, id string) (*domain.OrganisationD, error) {
	var model models.Organisation
	if err := r.db.
		WithContext(ctx).
		First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *OrganisationRepositoryImpl) GetByApiKey(ctx context.Context, apiKey string) (*domain.OrganisationD, error) {
	var model models.Organisation
	if err := r.db.
		WithContext(ctx).
		First(&model, "api_key = ?", apiKey).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *OrganisationRepositoryImpl) Register(ctx context.Context, request *domain.OrganisationD) (*domain.OrganisationD, error) {
	var model models.Organisation
	model.FromDomain(request)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}
