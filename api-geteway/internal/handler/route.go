package handler

import (
	"api-geteway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authHandler *AuthHandler,
	userHandler *UserHandler,
	categoryHandler *CategoryHandler,
	ingredientHandler *IngredientHandler,
	foodHandler *FoodHandler,
	nutritionHandler *,
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
	// ADMIN ROUTES
	// =========================
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// Users
		admin.POST("/users", userHandler.CreateUser)
		admin.GET("/users", userHandler.GetAllUsers)
		admin.GET("/users/:id", userHandler.GetUserByID)
		admin.PUT("/users/:id", userHandler.UpdateUser)
		admin.DELETE("/users/:id", userHandler.DeleteUser)

		// Categories
		admin.POST("/categories", categoryHandler.CreateCategory)
		admin.GET("/categories", categoryHandler.GetAllCategories)
		admin.GET("/categories/:id", categoryHandler.GetCategoryByID)
		admin.PUT("/categories/:id", categoryHandler.UpdateCategory)
		admin.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		// Restaurants
		admin.POST("/restaurants", foodHandler.CreateRestaurant)
		admin.GET("/restaurants", foodHandler.GetAllRestaurants)
		admin.GET("/restaurants/:id", foodHandler.GetRestaurantByID)
		admin.PUT("/restaurants/:id", foodHandler.UpdateRestaurant)
		admin.DELETE("/restaurants/:id", foodHandler.DeleteRestaurant)

		// Recipes
		admin.POST("/recipes", foodHandler.CreateRecipe)
		admin.GET("/recipes", foodHandler.GetAllRecipes)
		admin.GET("/recipes/:id", foodHandler.GetRecipeByID)
		admin.PUT("/recipes/:id", foodHandler.UpdateRecipe)
		admin.DELETE("/recipes/:id", foodHandler.DeleteRecipe)

		// Salads
admin.POST("/salads", foodHandler.CreateSalad)
admin.GET("/salads", foodHandler.GetAllSalads)       // ← qo'shildi
admin.GET("/salads/:id", foodHandler.GetSaladByID)
admin.PUT("/salads/:id", foodHandler.UpdateSalad)    // ← qo'shildi
admin.DELETE("/salads/:id", foodHandler.DeleteSalad) // ← qo'shildi
	}

	// =========================
	// USER ROUTES
	// =========================
	user := r.Group("")
	user.Use(middleware.AuthMiddleware())
	{
		// Categories (read only)
		user.GET("/categories", categoryHandler.GetAllCategories)
		user.GET("/categories/:id", categoryHandler.GetCategoryByID)
		user.GET("/categories/with-products", categoryHandler.GetAllWithUserProducts)

		// Ingredients (full CRUD)
		user.POST("/ingredients", ingredientHandler.CreateIngredient)
		user.GET("/ingredients", ingredientHandler.GetAllIngredients)
		user.GET("/ingredients/:id", ingredientHandler.GetIngredientByID)
		user.PUT("/ingredients/:id", ingredientHandler.UpdateIngredient)
		user.DELETE("/ingredients/:id", ingredientHandler.DeleteIngredient)

		// Food filter
		user.POST("/food/filter", foodHandler.FilterFood)

		// Statik 
		
	}

	return r
}