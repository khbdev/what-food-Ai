package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// =========================
// MIDDLEWARE
// =========================

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 🔥 secret ENV dan olinadi
		jwtSecret := []byte(os.Getenv("JWT_ACCESS_SECRET"))

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 🔥 expiry check
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// contextga user saqlash
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}