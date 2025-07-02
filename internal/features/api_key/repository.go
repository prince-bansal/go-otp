package api_key

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/models"
	"github.com/prince-bansal/go-otp/store/db"
	"gorm.io/gorm"
)

type ApiKeyRepository interface {
	Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiKeyD, error)
	Get(ctx context.Context, orgId int) ([]*domain.ApiKeyD, error)
	SoftDelete(ctx context.Context, d *domain.ApiKeyD) (*domain.ApiKeyD, error)
	FindById(ctx context.Context, id int) (*domain.ApiKeyD, error)
	GetBySaltHash(ctx context.Context, salt string) (*domain.ApiKeyD, error)
}

type Impl struct {
	db *gorm.DB
}

func NewApiKeyRepository(db *db.Store) ApiKeyRepository {
	return &Impl{
		db: db.Db,
	}
}

func (r *Impl) Create(ctx context.Context, request *domain.ApiKeyD) (*domain.ApiKeyD, error) {
	var model models.ApiKey

	model.FromDomain(request)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *Impl) Get(ctx context.Context, orgId int) ([]*domain.ApiKeyD, error) {
	var keys []*models.ApiKey
	if err := r.db.WithContext(ctx).
		Preload("Organisation").
		Find(&keys, "organisation_id = ?", orgId).Error; err != nil {
		return nil, err
	}

	var processedKeys []*domain.ApiKeyD

	for _, k := range keys {
		processedKeys = append(processedKeys, k.ToDomain())
	}
	return processedKeys, nil

}

func (r *Impl) SoftDelete(ctx context.Context, d *domain.ApiKeyD) (*domain.ApiKeyD, error) {
	var model models.ApiKey
	model.FromDomain(d)
	if err := r.db.WithContext(ctx).Delete(&model, d.Id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *Impl) FindById(ctx context.Context, id int) (*domain.ApiKeyD, error) {
	var model models.ApiKey
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}
func (r *Impl) GetBySaltHash(ctx context.Context, salt string) (*domain.ApiKeyD, error) {
	var model models.ApiKey
	if err := r.db.
		WithContext(ctx).
		Preload("Organisation").
		First(&model, "salt_hash = ?", salt).
		Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}
