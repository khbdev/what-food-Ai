package handler

import (
	"errors"
	"net/http"
	"strconv"

	"api-geteway/internal/models"
	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	"github.com/gin-gonic/gin"

	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

type FoodHandler struct {
	svc *service.FoodService
}

// =========================
// INIT
// =========================

func NewFoodHandler(s *service.FoodService) *FoodHandler {
	return &FoodHandler{svc: s}
}

// =========================
// RECIPE
// =========================

func (h *FoodHandler) CreateRecipe(c *gin.Context) {
	var req models.CreateRecipeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	err := h.svc.CreateRecipe(c.Request.Context(), &foodpb.CreateRecipeRequest{
		Recipe: &foodpb.Recipe{
			RestaurantId: req.Recipe.RestaurantID,
			Name:         req.Recipe.Name,
			Description:  req.Recipe.Description,
			ImageUrl:     req.Recipe.ImageURL,
			VideoUrl:     req.Recipe.VideoURL,
			Country:      req.Recipe.Country,
			MealTime:     req.Recipe.MealTime,
			Kcal:         req.Recipe.Kcal,
			Protein:      req.Recipe.Protein,
			Fat:          req.Recipe.Fat,
			Carbs:        req.Recipe.Carbs,
		},
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "recipe created"})
}

func (h *FoodHandler) GetRecipeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.GetRecipeByID(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

func (h *FoodHandler) GetAllRecipes(c *gin.Context) {
	res, err := h.svc.GetAllRecipes(c.Request.Context())
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

func (h *FoodHandler) UpdateRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req models.UpdateRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	req.Recipe.ID = id

	err = h.svc.UpdateRecipe(c.Request.Context(), &foodpb.UpdateRecipeRequest{
		Recipe: &foodpb.Recipe{
			Id:           req.Recipe.ID,
			RestaurantId: req.Recipe.RestaurantID,
			Name:         req.Recipe.Name,
			Description:  req.Recipe.Description,
			ImageUrl:     req.Recipe.ImageURL,
			VideoUrl:     req.Recipe.VideoURL,
			Country:      req.Recipe.Country,
			MealTime:     req.Recipe.MealTime,
			Kcal:         req.Recipe.Kcal,
			Protein:      req.Recipe.Protein,
			Fat:          req.Recipe.Fat,
			Carbs:        req.Recipe.Carbs,
		},
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "recipe updated"})
}

func (h *FoodHandler) DeleteRecipe(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = h.svc.DeleteRecipe(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "recipe deleted"})
}

// =========================
// SALAD
// =========================

func (h *FoodHandler) CreateSalad(c *gin.Context) {
	var req models.CreateSaladRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	err := h.svc.CreateSalad(c.Request.Context(), &foodpb.CreateSaladRequest{
		Salad: &foodpb.Salad{
			RestaurantId: req.Salad.RestaurantID,
			Name:         req.Salad.Name,
			Description:  req.Salad.Description,
			ImageUrl:     req.Salad.ImageURL,
			VideoUrl:     req.Salad.VideoURL,
			Country:      req.Salad.Country,
			MealTime:     req.Salad.MealTime,
			Kcal:         req.Salad.Kcal,
			Protein:      req.Salad.Protein,
			Fat:          req.Salad.Fat,
			Carbs:        req.Salad.Carbs,
		},
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "salad created"})
}

func (h *FoodHandler) GetSaladByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.GetSaladByID(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

func (h *FoodHandler) GetAllSalads(c *gin.Context) {
	res, err := h.svc.GetAllSalads(c.Request.Context())
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

func (h *FoodHandler) UpdateSalad(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req models.UpdateSaladRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	req.Salad.ID = id

	err = h.svc.UpdateSalad(c.Request.Context(), &foodpb.UpdateSaladRequest{
		Salad: &foodpb.Salad{
			Id:           req.Salad.ID,
			RestaurantId: req.Salad.RestaurantID,
			Name:         req.Salad.Name,
			Description:  req.Salad.Description,
			ImageUrl:     req.Salad.ImageURL,
			VideoUrl:     req.Salad.VideoURL,
			Country:      req.Salad.Country,
			MealTime:     req.Salad.MealTime,
			Kcal:         req.Salad.Kcal,
			Protein:      req.Salad.Protein,
			Fat:          req.Salad.Fat,
			Carbs:        req.Salad.Carbs,
		},
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "salad updated"})
}

func (h *FoodHandler) DeleteSalad(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = h.svc.DeleteSalad(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "salad deleted"})
}
// =========================
// RESTAURANT
// =========================

func (h *FoodHandler) CreateRestaurant(c *gin.Context) {
	var req models.CreateRestaurantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.svc.CreateRestaurant(c.Request.Context(), &foodpb.CreateRestaurantRequest{
		RestaurantName: req.RestaurantName,
		Description:    req.Description,
		ImageUrl:       req.ImageURL,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, models.CreateRestaurantResponse{ID: id})
}

func (h *FoodHandler) GetRestaurantByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.GetRestaurantByID(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

func (h *FoodHandler) GetAllRestaurants(c *gin.Context) {
	res, err := h.svc.GetAllRestaurants(c.Request.Context())
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}
func (h *FoodHandler) UpdateRestaurant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req models.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	err = h.svc.UpdateRestaurant(c.Request.Context(), &foodpb.UpdateRestaurantRequest{
		Restaurant: &foodpb.Restaurant{
			Id:             id, // 🔥 MUHIM FIX
			RestaurantName: req.RestaurantName,
			Description:    req.Description,
			ImageUrl:       req.ImageURL,
		},
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "restaurant updated"})
}
func (h *FoodHandler) DeleteRestaurant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = h.svc.DeleteRestaurant(c.Request.Context(), &foodpb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, gin.H{"message": "restaurant deleted"})
}

// =========================
// FILTER
// =========================

func (h *FoodHandler) FilterFood(c *gin.Context) {
	var req models.FoodFilterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.FilterFood(c.Request.Context(), &foodpb.FoodFilterRequest{
		Country:       req.Country,
		MealTime:      req.MealTime,
		MaxKcal:       req.MaxKcal,
		IncludeSalads: req.IncludeSalads,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}