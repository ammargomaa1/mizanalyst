package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/responses"
	"github.com/mizanalyst/mizanalyst/utils"
)

// AuthMiddleware validates the Bearer token from the Authorization header,
// parses the access token claims, and injects user_id into the gin context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Unauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			responses.Unauthorized(c, "Authorization header must be in the format: Bearer <token>")
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := utils.ParseAccessToken(tokenString)
		if err != nil {
			responses.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
