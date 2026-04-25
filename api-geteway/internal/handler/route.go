package handler

import (
	
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *AuthHandler, userHandler ) *gin.Engine {

	r := gin.Default()

	// =========================
	// AUTH ROUTES
	// =========================
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/verify", authHandler.VerifyOTP)
	}

	return r
}