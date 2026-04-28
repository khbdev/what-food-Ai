package handler

import (
	"github.com/gin-gonic/gin"
	"api-geteway/internal/middleware"
)

func SetupRouter(
	authHandler *AuthHandler,
	userHandler *UserHandler,
	categoryHandler *CategoryHandler,
	ingredientHandler *IngredientHandler,
) *gin.Engine {

	r := gin.Default()

	// =========================
	// AUTH (PUBLIC)
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
		admin.PUT("/user/:id", userHandler.UpdateUser)
		admin.DELETE("/user/:id", userHandler.DeleteUser)
	}

	// =========================
	// CATEGORY ROUTES
	// =========================
	category := r.Group("/categories")
	category.Use(middleware.AuthMiddleware())
	{
		// READ (USER + ADMIN)
		category.GET("", categoryHandler.GetAllCategories)
		category.GET("/:id", categoryHandler.GetCategoryByID)
		category.GET("/with-products", categoryHandler.GetAllWithUserProducts)

		// WRITE (ADMIN ONLY)
		adminCat := category.Group("")
		adminCat.Use(middleware.AdminMiddleware())
		{
			adminCat.POST("", categoryHandler.CreateCategory)
			adminCat.PUT("/:id", categoryHandler.UpdateCategory)
			adminCat.DELETE("/:id", categoryHandler.DeleteCategory)
		}
	}

	// =========================
	// INGREDIENT ROUTES (USER ONLY)
	// =========================
	ingredients := r.Group("/ingredients")
	ingredients.Use(middleware.AuthMiddleware())
	{
		ingredients.POST("", ingredientHandler.CreateIngredient)
		ingredients.GET("", ingredientHandler.GetAllIngredients)
		ingredients.GET("/:id", ingredientHandler.GetIngredientByID)
		ingredients.PUT("/:id", ingredientHandler.UpdateIngredient)
		ingredients.DELETE("/:id", ingredientHandler.DeleteIngredient)
	}

	return r
}