package middleware

import (
	"net/http"
	"strings"

	"go-auth-api/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		username, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan username ke context
		c.Set("username", username)
		c.Next()
	}
}
