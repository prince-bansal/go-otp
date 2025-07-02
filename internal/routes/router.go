package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/api_key"
	"github.com/prince-bansal/go-otp/internal/features/health"
	"github.com/prince-bansal/go-otp/internal/features/organisation"
	"github.com/prince-bansal/go-otp/internal/features/otp"
)

type Router struct {
	organisationHandler *organisation.OrganisationHandler
	healthHandler       *health.HealthHandler
	apiKeyHandler       *api_key.ApiKeyHandler
	otpHandler          *otp.OtpHandler
}

func NewRouter(handler *organisation.OrganisationHandler, healthHandler *health.HealthHandler, keyHandler *api_key.ApiKeyHandler, otpHandler *otp.OtpHandler) *Router {
	return &Router{
		organisationHandler: handler,
		healthHandler:       healthHandler,
		apiKeyHandler:       keyHandler,
		otpHandler:          otpHandler,
	}
}

func (r *Router) InitRoutes(router *gin.Engine) {
	r.healthHandler.InitRoutes(router)
	r.organisationHandler.InitRoutes(router)
	r.apiKeyHandler.InitRoutes(router)
	r.otpHandler.InitRoutes(router)
}
