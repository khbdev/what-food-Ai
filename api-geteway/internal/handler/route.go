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
	foodHandler *FoodHandler,
	aiFoodHandler *mailhandler.FoodHandler,
	nutritionHandler *NutritionHandler,
) *gin.Engine {

	r := gin.New()

	// =========================
	// GLOBAL MIDDLEWARE
	// =========================
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

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
		// USERS
		admin.POST("/users", userHandler.CreateUser)
		admin.GET("/users", userHandler.GetAllUsers)
		admin.GET("/users/:id", userHandler.GetUserByID)
		admin.PUT("/users/:id", userHandler.UpdateUser)
		admin.DELETE("/users/:id", userHandler.DeleteUser)

		// CATEGORIES
		admin.POST("/categories", categoryHandler.CreateCategory)
		admin.GET("/categories", categoryHandler.GetAllCategories)
		admin.GET("/categories/:id", categoryHandler.GetCategoryByID)
		admin.PUT("/categories/:id", categoryHandler.UpdateCategory)
		admin.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		// RESTAURANTS
		admin.POST("/restaurants", foodHandler.CreateRestaurant)
		admin.GET("/restaurants", foodHandler.GetAllRestaurants)
		admin.GET("/restaurants/:id", foodHandler.GetRestaurantByID)
		admin.PUT("/restaurants/:id", foodHandler.UpdateRestaurant)
		admin.DELETE("/restaurants/:id", foodHandler.DeleteRestaurant)

		// RECIPES
		admin.POST("/recipes", foodHandler.CreateRecipe)
		admin.GET("/recipes", foodHandler.GetAllRecipes)
		admin.GET("/recipes/:id", foodHandler.GetRecipeByID)
		admin.PUT("/recipes/:id", foodHandler.UpdateRecipe)
		admin.DELETE("/recipes/:id", foodHandler.DeleteRecipe)

		// SALADS
		admin.POST("/salads", foodHandler.CreateSalad)
		admin.GET("/salads", foodHandler.GetAllSalads)
		admin.GET("/salads/:id", foodHandler.GetSaladByID)
		admin.PUT("/salads/:id", foodHandler.UpdateSalad)
		admin.DELETE("/salads/:id", foodHandler.DeleteSalad)
	}

	// =========================
	// USER ROUTES (AUTH REQUIRED)
	// =========================
	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		// CATEGORIES
		user.GET("/categories", categoryHandler.GetAllCategories)
		user.GET("/categories/:id", categoryHandler.GetCategoryByID)
		user.GET("/categories/with-products", categoryHandler.GetAllWithUserProducts)

		// INGREDIENTS
		user.POST("/ingredients", ingredientHandler.CreateIngredient)
		user.GET("/ingredients", ingredientHandler.GetAllIngredients)
		user.GET("/ingredients/:id", ingredientHandler.GetIngredientByID)
		user.PUT("/ingredients/:id", ingredientHandler.UpdateIngredient)
		user.DELETE("/ingredients/:id", ingredientHandler.DeleteIngredient)

		// AI FOOD
		user.POST("/food/filter", aiFoodHandler.FilterFood)
		user.GET("/food/detail", aiFoodHandler.GetFoodDetail)

		// Statik
		user.GET("/nutrition/weekly", nutritionHandler.GetWeeklyNutrition)
	}

	return r
}