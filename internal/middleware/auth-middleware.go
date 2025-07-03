package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/domain/response"
	"github.com/prince-bansal/go-otp/internal/features/api_key"
	"github.com/prince-bansal/go-otp/internal/utils/constants"
	"github.com/prince-bansal/go-otp/pkg/logger"
)

type Middleware struct {
	apiService api_key.ApiService
}

func NewMiddleware(apiService api_key.ApiService) *Middleware {
	return &Middleware{
		apiService: apiService,
	}
}

func (m *Middleware) ApiGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(constants.API_KEY)
		if key == "" {
			logger.Error("api key is missing")
			c.JSON(401, response.SendAuthenticationError())
			c.Abort()
			return
		}
		org, err := m.apiService.GetByApiKey(c, key)
		if err != nil {
			logger.Error("middleware", "could not find record by api key", err)
			c.JSON(401, response.SendInvalidError("invalid request", err))
			c.Abort()
			return

		}
		c.Set("OrganisationId", org.OrganisationId)
		c.Next()
	}
}
