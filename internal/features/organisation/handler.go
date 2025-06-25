package organisation

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
)

type OrganisationHandler struct {
	organisationService OrganisationService
}

func NewOrganisationHandler(organisationService OrganisationService) *OrganisationHandler {
	return &OrganisationHandler{
		organisationService: organisationService,
	}
}

func (h *OrganisationHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/organisation")
	{
		routes.GET("/", h.getAll)
		routes.GET("/:id", h.getOne)
		routes.POST("/", h.createOrganisation)
	}

}

func (h *OrganisationHandler) createOrganisation(ctx *gin.Context) {
	var req domain.OrganisationRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":     "Invalid Request",
			"errorType": "Validation Error",
		})
		return
	}

	records := h.organisationService.Register(ctx, req)
	ctx.JSON(201, records)
}

func (h *OrganisationHandler) getAll(ctx *gin.Context) {
	records := h.organisationService.GetAll(ctx)
	ctx.JSON(201, records)
}

func (h *OrganisationHandler) getOne(ctx *gin.Context) {
	// todo: validate id
	id := ctx.Param("id")
	records := h.organisationService.GetOne(ctx, id)
	ctx.JSON(201, records)
}
