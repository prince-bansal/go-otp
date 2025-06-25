package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/health"
	"github.com/prince-bansal/go-otp/internal/features/organisation"
)

type Router struct {
	organisationHandler *organisation.OrganisationHandler
	healthHandler       *health.HealthHandler
}

func NewRouter(handler *organisation.OrganisationHandler, healthHandler *health.HealthHandler) *Router {
	return &Router{
		organisationHandler: handler,
		healthHandler:       healthHandler,
	}
}

func (r *Router) InitRoutes(router *gin.Engine) {
	r.healthHandler.InitRoutes(router)
	r.organisationHandler.InitRoutes(router)
}
