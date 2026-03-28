package middleware

import (
	"net/http"
	"sppg-backend/internal/model"
	"sppg-backend/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Message: "Access token is required",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Message: "Invalid token format, use: Bearer <token>",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Message: "Token is invalid or expired",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		allowed := false
		for _, r := range roles {
			if r == role {
				allowed = true
				break
			}
		}
		if !allowed {
			c.JSON(http.StatusForbidden, model.Response{
				Success: false,
				Message: "You don't have permission to access this resource",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}