package organisation

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/organisation/domain"
	"github.com/prince-bansal/go-otp/internal/utils"
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
		ctx.JSON(400, Response.SendValidationError(err))
		return
	}

	err = req.Validate()
	if err != nil {
		ctx.JSON(400, Response.SendValidationError(err))
		return
	}

	data := &domain.OrganisationD{
		Name:  req.Name,
		Email: req.Email,
	}

	record, err := h.organisationService.Register(ctx, data)
	if err != nil {
		ctx.JSON(400, Response.SendError("getting error creating organisation", err))
		return
	}
	ctx.JSON(201, record)
}

func (h *OrganisationHandler) getAll(ctx *gin.Context) {
	records, err := h.organisationService.GetAll(ctx)
	if err != nil {
		ctx.JSON(400, Response.SendError("getting error creating organisation", err))
		return
	}
	ctx.JSON(200, records)
}

func (h *OrganisationHandler) getOne(ctx *gin.Context) {
	// todo: validate id
	id := ctx.Param("id")
	record, err := h.organisationService.GetOne(ctx, id)
	if err != nil {
		ctx.JSON(400, Response.SendError("cannot find organisation with id", err))
		return
	}
	ctx.JSON(200, record)
}
