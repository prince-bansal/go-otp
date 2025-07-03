package api_key

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/domain"
	"github.com/prince-bansal/go-otp/internal/domain/response"
	"github.com/prince-bansal/go-otp/internal/utils/timeutil"
	"github.com/prince-bansal/go-otp/pkg/logger"
	"strconv"
)

type ApiKeyHandler struct {
	apiService ApiService
}

func NewApiHandler(service ApiService) *ApiKeyHandler {
	return &ApiKeyHandler{
		apiService: service,
	}
}

func (h *ApiKeyHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/apikey")
	{
		routes.GET("/:orgId", h.getAllApiKeyByOrganisation)
		routes.DELETE("/:id", h.expireKey)
		routes.POST("/", h.createApiKey)
	}
}

func (h *ApiKeyHandler) createApiKey(ctx *gin.Context) {
	var req domain.ApiKeyRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error("failed to parse request body", err)
		ctx.JSON(400, response.SendInvalidError("unable to marshal request", err))
		return
	}

	err = req.Validate()
	if err != nil {
		logger.Error("invalid body", err)
		ctx.JSON(400, response.SendValidationError(err))
		return
	}

	parsedTime, err := timeutil.ConvertInYYYYMMDD(req.Expiry)
	if err != nil {
		logger.Error("failed to parse time", err)
		ctx.JSON(400, response.SendError("time conversion error", err))
		return
	}

	d := domain.ApiKeyD{
		OrganisationId: req.OrganisationId,
		Expiry:         parsedTime,
	}

	createdKey, err := h.apiService.Create(ctx, &d)
	if err != nil {
		logger.Error("failed to save api key", err)
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(201, response.SuccessWithStatus(201, createdKey))
	return
}

func (h *ApiKeyHandler) getAllApiKeyByOrganisation(ctx *gin.Context) {
	orgId := ctx.Param("orgId")
	intId, err := strconv.Atoi(orgId)
	if err != nil {
		logger.Error("failed to convert %s to int", orgId, err)
		ctx.JSON(400, err)
		return
	}
	apiKeys, err := h.apiService.GetAll(ctx, intId)
	if err != nil {
		logger.Error("failed to fetch api keys for %d", intId, err)
		ctx.JSON(400, err)
		return
	}
	ctx.JSON(200, response.Success(apiKeys))
	return
}

func (h *ApiKeyHandler) expireKey(ctx *gin.Context) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("failed to convert %s to int", id, err)
		ctx.JSON(400, err)
		return
	}

	deletedRecord, err := h.apiService.Expire(ctx, intId)
	if err != nil {
		logger.Error("failed to delete the api key %d", id, err)
		ctx.JSON(400, response.SendError("invalid request", err))
		return
	}
	ctx.JSON(200, response.Success(deletedRecord))
	return

}
