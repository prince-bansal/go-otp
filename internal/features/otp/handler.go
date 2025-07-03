package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/domain"
	response "github.com/prince-bansal/go-otp/internal/domain/response"
	"github.com/prince-bansal/go-otp/internal/middleware"
	"github.com/prince-bansal/go-otp/pkg/logger"
)

type OtpHandler struct {
	service OtpService
	m       *middleware.Middleware
}

func (h *OtpHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/otp")
	routes.DELETE("/", h.deleteExpiredOtp)
	{
		protectedRoutes := routes.Group("").Use(h.m.ApiGuard())
		{
			protectedRoutes.POST("/", h.generateOtp)
			protectedRoutes.POST("/verify", h.verifyOtp)
		}
	}
}

func NewOtpHandler(service OtpService, m *middleware.Middleware) *OtpHandler {
	return &OtpHandler{
		service: service,
		m:       m,
	}

}

func (h *OtpHandler) generateOtp(ctx *gin.Context) {

	var req domain.OtpGenerateRequest
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		logger.Error("failed to parse body", err)
		ctx.JSON(400, response.SendValidationError(err))
		return
	}

	otp, err := h.service.GenerateOtp(ctx, &req)

	if err != nil {
		logger.Error("failed to generate otp", err)
		ctx.JSON(400, response.SendInvalidError("cannot generate otp right now", err))
		return
	}
	ctx.JSON(200, otp)
	return
}

func (h *OtpHandler) verifyOtp(ctx *gin.Context) {

	var req domain.OtpVerifyRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error("failed to parse body", err)
		ctx.JSON(400, response.SendValidationError(err))
		return
	}

	success, err := h.service.VerifyOtp(ctx, &req)
	if err != nil || !success {
		logger.Error("invalid credentials", err)
		ctx.JSON(400, response.SendInvalidError("could not verify", err))
		return
	}

	ctx.JSON(200, response.Success("verified successfully"))
	return
}

func (h *OtpHandler) deleteExpiredOtp(ctx *gin.Context) {
	success, err := h.service.CleanOtps(ctx)
	if err != nil {
		logger.Error("failed to delete otps", err)
		ctx.JSON(200, response.SendInvalidError("could not delete otps", err))
		return
	}
	ctx.JSON(200, response.Success(success))
}
