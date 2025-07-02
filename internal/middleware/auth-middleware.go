package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-bansal/go-otp/internal/domain/response"
	"github.com/prince-bansal/go-otp/internal/features/apiKey"
	"github.com/prince-bansal/go-otp/internal/utils/constants"
)

type Middleware struct {
	apiService apiKey.ApiService
}

func NewMiddleware(apiService apiKey.ApiService) *Middleware {
	return &Middleware{
		apiService: apiService,
	}
}

func (m *Middleware) ApiGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(constants.API_KEY)
		if key == "" {
			c.JSON(401, response.SendAuthenticationError())
			c.Abort()
			return
		}
		org, err := m.apiService.GetByApiKey(c, key)
		if err != nil {
			c.JSON(401, response.SendInvalidError("invalid request", err))
			c.Abort()
			return

		}
		c.Set("OrganisationId", org.OrganisationId)
		c.Next()
	}
}
