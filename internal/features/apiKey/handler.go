package apiKey

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/apiKey/domain"
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
		routes.POST("/", h.createApiKey)
	}
}

func (h *ApiKeyHandler) createApiKey(ctx *gin.Context) {
	var req domain.ApiKeyRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	createdKey := h.apiService.Create(ctx, req)
	ctx.JSON(201, createdKey)
	return
}

func (h *ApiKeyHandler) getAllApiKeyByOrganisation(ctx *gin.Context) {
	orgId := ctx.Param("id")
	apiKeys := h.apiService.GetAll(ctx, orgId)
	ctx.JSON(200, apiKeys)
	return
}
