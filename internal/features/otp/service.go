package otp

import (
	"context"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/internal/features/apiKey"
)

type OtpService interface {
	GenerateOtp(ctx context.Context, otp *domain.OtpGenerateRequest) (*domain.OtpGenerateResponse, error)
	VerifyOtp(ctx context.Context, request *domain.OtpVerifyRequest) (bool, error)
	CleanOtps(ctx context.Context) (bool, error)
}

type impl struct {
	repository OtpRepository
	apiService apiKey.ApiService
}

func NewOtpService(repository OtpRepository, api apiKey.ApiService) OtpService {
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

	return s.repository.Verify(ctx, &d)
}

func (s *impl) CleanOtps(ctx context.Context) (bool, error) {
	return s.repository.DeleteExpired(ctx)
}
