package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://localhost:5173",
			"https://monumental-pithivier-264317.netlify.app",
		}

		origin := c.Request.Header.Get("Origin")
		var allowedOrigin string
		for _, o := range allowedOrigins {
			if o == origin {
				allowedOrigin = o
				break
			}
		}

		if allowedOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Writer.Header().Add("Vary", "Origin")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}