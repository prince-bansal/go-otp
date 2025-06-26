package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/apiKey"
	"github.com/prince-bansal/go-otp/internal/features/health"
	"github.com/prince-bansal/go-otp/internal/features/organisation"
)

type Router struct {
	organisationHandler *organisation.OrganisationHandler
	healthHandler       *health.HealthHandler
	apiKeyHandler       *apiKey.ApiKeyHandler
}

func NewRouter(handler *organisation.OrganisationHandler, healthHandler *health.HealthHandler, keyHandler *apiKey.ApiKeyHandler) *Router {
	return &Router{
		organisationHandler: handler,
		healthHandler:       healthHandler,
		apiKeyHandler:       keyHandler,
	}
}

func (r *Router) InitRoutes(router *gin.Engine) {
	r.healthHandler.InitRoutes(router)
	r.organisationHandler.InitRoutes(router)
	r.apiKeyHandler.InitRoutes(router)
}
