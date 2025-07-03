package otp

import (
	"context"
	"errors"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/models"
	"github.com/prince-bansal/go-otp/pkg/logger"
	"github.com/prince-bansal/go-otp/store/db"
	"gorm.io/gorm"
	"time"
)

type OtpRepository interface {
	Insert(ctx context.Context, otp *domain.Otp) (*domain.Otp, error)
	DeleteExpired(ctx context.Context) (bool, error)
	Verify(ctx context.Context, request *domain.Otp) (bool, error)
}

type Impl struct {
	db *gorm.DB
}

func NewOtpRepository(db *db.Store) OtpRepository {
	return &Impl{
		db: db.Db,
	}
}

func (r *Impl) Insert(ctx context.Context, domain *domain.Otp) (*domain.Otp, error) {
	var model models.Otp
	model.FromDomain(domain)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		logger.Error("failed to save otp", err)
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *Impl) DeleteExpired(ctx context.Context) (bool, error) {
	var model models.Otp
	endTime := time.Now()

	if err := r.db.
		WithContext(ctx).
		Delete(&model, "created_at < ?", endTime).Error; err != nil {
		logger.Error("failed to delete expired otps", err)
		return false, err
	}

	return true, nil
}

func (r *Impl) Verify(ctx context.Context, domain *domain.Otp) (bool, error) {
	var model models.Otp
	model.FromDomain(domain)

	nMinutesAgo := time.Now().Add(-10 * time.Minute)

	deletedRows := r.db.
		WithContext(ctx).
		Delete(&model, "mobile_number = ? AND otp = ? AND created_at > ?", domain.MobileNo, domain.Otp, nMinutesAgo)

	if deletedRows.Error != nil {
		logger.Error("otp not found for number %d", domain.MobileNo)
		return false, deletedRows.Error
	}

	if deletedRows.RowsAffected == 0 {
		return false, errors.New("invalid otp")
	}
	return true, nil
}
