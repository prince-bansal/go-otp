package middleware

import (
	"github.com/gin-gonic/gin"
	Response "github.com/prince-bansal/go-otp/internal/utils"
)

func ApiGuard() gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("API_KEY")
		if apiKey == "" {
			c.JSON(401, Response.SendAuthenticationError())
			c.Abort()
			return
		}
		c.Set("API_KEY", apiKey)
		c.Next()
	}
}
