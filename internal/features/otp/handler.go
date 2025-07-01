package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/features/otp/domain"
	Response "github.com/prince-bansal/go-otp/internal/utils"
)

type OtpHandler struct {
	service OtpService
}

func (h *OtpHandler) InitRoutes(router *gin.Engine) {
	routes := router.Group("/otp")
	{
		routes.POST("/", h.generateOtp)
		routes.POST("/verify", h.verifyOtp)
		routes.DELETE("/", h.deleteExpiredOtp)
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

	apiKey := ctx.GetHeader("API_KEY")

	if apiKey == "" {
		ctx.JSON(400, Response.SendAuthenticationError())
		return
	}

	otp, err := h.service.GenerateOtp(ctx, &req, apiKey)

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

	apiKey := ctx.GetHeader("API_KEY")
	if apiKey == "" {
		ctx.JSON(400, Response.SendAuthenticationError())
		return
	}

	success, err := h.service.VerifyOtp(ctx, &req, apiKey)
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
