package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/otp/domain"
	"github.com/prince-bansal/go-otp/internal/middleware"
	Response "github.com/prince-bansal/go-otp/internal/utils"
)

type OtpHandler struct {
	service OtpService
}

func (h *OtpHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/otp")
	routes.DELETE("/", h.deleteExpiredOtp)
	{
		protectedRoutes := routes.Group("").Use(middleware.ApiGuard())
		{
			protectedRoutes.POST("/", h.generateOtp)
			protectedRoutes.POST("/verify", h.verifyOtp)
		}
	}
}

func NewOtpHandler(service OtpService) *OtpHandler {
	return &OtpHandler{
		service: service,
	}
}

func (h *OtpHandler) generateOtp(ctx *gin.Context) {

	var req domain.OtpGenerateRequest
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(400, Response.SendValidationError(err))
		return
	}

	apiKey, _ := ctx.Get("API_KEY")

	otp, err := h.service.GenerateOtp(ctx, &req, apiKey.(string))

	if err != nil {
		ctx.JSON(400, Response.SendInvalidError("cannot generate otp right now", err))
		return
	}
	ctx.JSON(200, otp)
	return
}

func (h *OtpHandler) verifyOtp(ctx *gin.Context) {

	var req domain.OtpVerifyRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, Response.SendValidationError(err))
		return
	}

	apiKey, _ := ctx.Get("API_KEY")

	success, err := h.service.VerifyOtp(ctx, &req, apiKey.(string))
	if err != nil || !success {
		ctx.JSON(400, Response.SendInvalidError("could not verify", err))
		return
	}

	ctx.JSON(200, Response.Success("verified successfully"))
	return
}

func (h *OtpHandler) deleteExpiredOtp(ctx *gin.Context) {
	success, err := h.service.CleanOtps(ctx)
	if err != nil {
		ctx.JSON(200, Response.SendInvalidError("could not delete otps", err))
		return
	}
	ctx.JSON(200, Response.Success(success))
}
