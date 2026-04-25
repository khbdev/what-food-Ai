package handler

import (
	"github.com/gin-gonic/gin"
	"api-geteway/internal/middleware"
)

func SetupRouter(authHandler *AuthHandler, userHandler *UserHandler) *gin.Engine {

	r := gin.Default()

	// =========================
	// AUTH ROUTES (public)
	// =========================
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/verify", authHandler.VerifyOTP)
	}



	// =========================
	// ADMIN ROUTES (ONLY ADMIN)
	// =========================
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/user", userHandler.CreateUser)
		admin.GET("/user/all", userHandler.GetAllUsers)
		admin.GET("/user/:id", userHandler.GetUserByID)
		admin.PUT("/user", userHandler.UpdateUser)
		admin.DELETE("/user", userHandler.DeleteUser)
	}

	return r
}