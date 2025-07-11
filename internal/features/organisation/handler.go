package organisation

import (
	"github.com/gin-gonic/gin"
	domain "github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/internal/domain/response"
	"github.com/prince-bansal/go-otp/pkg/logger"
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
		logger.Error("failed to parse request body", err)
		ctx.JSON(400, response.SendValidationError(err))
		return
	}

	err = req.Validate()
	if err != nil {
		logger.Error("invalid body", err)
		ctx.JSON(400, response.SendValidationError(err))
		return
	}

	data := &domain.OrganisationD{
		Name:  req.Name,
		Email: req.Email,
	}

	record, err := h.organisationService.Register(ctx, data)
	if err != nil {
		logger.Error("failed to save organisation", err)
		ctx.JSON(400, response.SendError("getting error creating organisation", err))
		return
	}
	ctx.JSON(201, response.Success(record))
}

func (h *OrganisationHandler) getAll(ctx *gin.Context) {
	records, err := h.organisationService.GetAll(ctx)
	if err != nil {
		logger.Error("failed to fetch organisations", err)
		ctx.JSON(400, response.SendError("getting error creating organisation", err))
		return
	}
	ctx.JSON(200, response.Success(records))
}

func (h *OrganisationHandler) getOne(ctx *gin.Context) {
	// todo: validate id
	id := ctx.Param("id")
	record, err := h.organisationService.GetOne(ctx, id)
	if err != nil {
		logger.Error("organisation not found: %s", id, err)
		ctx.JSON(400, response.SendError("cannot find organisation with id", err))
		return
	}
	ctx.JSON(200, response.Success(record))
}
