package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// getAllowedOrigins — environment variable yoki default qiymatlardan originlarni qaytaradi.
// ALLOWED_ORIGINS=http://localhost:3000,https://yourapp.com ko'rinishida set qiling.
func getAllowedOrigins() map[string]bool {
	origins := map[string]bool{}

	raw := os.Getenv("ALLOWED_ORIGINS")
	if raw == "" {
		// default: faqat local development
		raw = "http://localhost:3000"
	}

	for _, o := range strings.Split(raw, ",") {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" {
			origins[trimmed] = true
		}
	}

	return origins
}

// CORSMiddleware — Cross-Origin Resource Sharing headerlarini boshqaradi.
// Faqat ruxsat etilgan originlar uchun CORS headerlar o'rnatadi.
// OPTIONS preflight requestlarini 204 bilan qaytaradi.
func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := getAllowedOrigins()

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if origin != "" && allowedOrigins[origin] {
			header := c.Writer.Header()
			header.Set("Access-Control-Allow-Origin", origin)
			header.Set("Access-Control-Allow-Credentials", "true")
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			header.Set("Access-Control-Max-Age", "86400") // 24 soat — preflight cache
			header.Add("Vary", "Origin")                  // proxy cache uchun muhim
		}

		// Preflight request — header o'rnatilgandan keyin tugatamiz
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // 204
			return
		}

		c.Next()
	}
}