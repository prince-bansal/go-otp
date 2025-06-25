package health

import "github.com/gin-gonic/gin"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/health")
	{
		routes.GET("/health", h.checkHealth)
	}
}

func (h *HealthHandler) checkHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"HEALTH": "UP",
	})
}
