package handler

import (
	mailhandler "api-geteway/internal/handler/mail_handler"
	"api-geteway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authHandler *AuthHandler,
	userHandler *UserHandler,
	categoryHandler *CategoryHandler,
	ingredientHandler *IngredientHandler,

	// OLD FOOD CRUD
	foodHandler *FoodHandler,

	// NEW AI FOOD HANDLER
	aiFoodHandler *mailhandler.FoodHandler,

	nutritionHandler *NutritionHandler,
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
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.AdminMiddleware(),
	)

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
		admin.GET("/salads", foodHandler.GetAllSalads)
		admin.GET("/salads/:id", foodHandler.GetSaladByID)
		admin.PUT("/salads/:id", foodHandler.UpdateSalad)
		admin.DELETE("/salads/:id", foodHandler.DeleteSalad)
	}

	// =========================
	// USER ROUTES
	// =========================

	user := r.Group("")
	user.Use(middleware.AuthMiddleware())

	{
		// Categories
		user.GET("/categories", categoryHandler.GetAllCategories)
		user.GET("/categories/:id", categoryHandler.GetCategoryByID)
		user.GET("/categories/with-products", categoryHandler.GetAllWithUserProducts)

		// Ingredients
		user.POST("/ingredients", ingredientHandler.CreateIngredient)
		user.GET("/ingredients", ingredientHandler.GetAllIngredients)
		user.GET("/ingredients/:id", ingredientHandler.GetIngredientByID)
		user.PUT("/ingredients/:id", ingredientHandler.UpdateIngredient)
		user.DELETE("/ingredients/:id", ingredientHandler.DeleteIngredient)

		// =========================
		// AI FOOD
		// =========================

		user.POST("/food/filter", aiFoodHandler.FilterFood)
		user.GET("/food/detail", aiFoodHandler.GetFoodDetail)

		// =========================
		// NUTRITION
		// =========================

		user.POST("/nutrition/weekly", nutritionHandler.GetWeeklyNutrition)
	}

	return r
}