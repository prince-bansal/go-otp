package otp

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/internal/features/api_key"
	"github.com/prince-bansal/go-otp/pkg/logger"
)

type OtpService interface {
	GenerateOtp(ctx context.Context, otp *domain.OtpGenerateRequest) (*domain.OtpGenerateResponse, error)
	VerifyOtp(ctx context.Context, request *domain.OtpVerifyRequest) (bool, error)
	CleanOtps(ctx context.Context) (bool, error)
}

type impl struct {
	repository OtpRepository
	apiService api_key.ApiService
}

func NewOtpService(repository OtpRepository, api api_key.ApiService) OtpService {
	return &impl{
		repository: repository,
		apiService: api,
	}
}

func (s *impl) GenerateOtp(ctx context.Context, req *domain.OtpGenerateRequest) (*domain.OtpGenerateResponse, error) {
	orgId := ctx.Value("OrganisationId")

	d := domain.Otp{
		OrganisationId: orgId.(int),
		MobileNo:       req.MobileNo,
	}
	d.Otp = d.GenerateOtp()
	record, err := s.repository.Insert(ctx, &d)
	if err != nil {
		logger.Error("failed to save otp", err)
		return nil, err
	}

	return &domain.OtpGenerateResponse{
		Otp: record.Otp,
	}, err
}

func (s *impl) VerifyOtp(ctx context.Context, req *domain.OtpVerifyRequest) (bool, error) {
	orgId := ctx.Value("OrganisationId")
	d := domain.Otp{
		Otp:            req.Otp,
		OrganisationId: orgId.(int),
		MobileNo:       req.MobileNo,
	}

	success, err := s.repository.Verify(ctx, &d)
	if err != nil {
		logger.Error("invalid credentials", err)
		return false, nil
	}
	return success, nil
}

func (s *impl) CleanOtps(ctx context.Context) (bool, error) {
	success, err := s.repository.DeleteExpired(ctx)
	if err != nil {
		logger.Error("failed to delete expired otps", err)
		return false, nil
	}
	return success, nil
}
