package mailhandler

import (
	"net/http"
	"strconv"

	mailmodels "api-geteway/internal/models/mailmodels"
	

	"github.com/gin-gonic/gin"

	asosiypb "github.com/khbdev/what-food-proto/proto/asosiy"
)

type FoodHandler struct {
	service *mail.FoodService
}

// =========================
// INIT
// =========================

func NewFoodHandler(service *mail.FoodService) *FoodHandler {
	return &FoodHandler{
		service: service,
	}
}

// =========================
// FILTER FOOD
// =========================

func (h *FoodHandler) FilterFood(c *gin.Context) {

	kcalLimit, err := strconv.Atoi(c.DefaultQuery("kcal_limit", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid kcal_limit",
		})
		return
	}

	hasSalad, err := strconv.ParseBool(c.DefaultQuery("has_salad", "false"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid has_salad",
		})
		return
	}

	req := &asosiypb.FoodFilterRequest{
		Country:   c.Query("country"),
		MealTime:  c.Query("meal_time"),
		HasSalad:  hasSalad,
		KcalLimit: int32(kcalLimit),
	}

	res, err := h.service.FilterFood(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	foods := make([]mailmodels.FoodItem, 0)

	for _, item := range res.Items {

		foods = append(foods, mailmodels.FoodItem{
			ID:           item.Id,
			Type:         item.Type,
			RestaurantID: item.RestaurantId,
			Name:         item.Name,
			Description:  item.Description,
			ImageURL:     item.ImageUrl,
			VideoURL:     item.VideoUrl,
			Country:      item.Country,
			MealTime:     item.MealTime,
			Kcal:         item.Kcal,
			Protein:      item.Protein,
			Fat:          item.Fat,
			Carbs:        item.Carbs,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(foods),
		"data":    foods,
	})
}

// =========================
// GET FOOD DETAIL
// =========================

func (h *FoodHandler) GetFoodDetail(c *gin.Context) {

	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	portion, err := strconv.Atoi(c.DefaultQuery("portion", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid portion",
		})
		return
	}

	req := &asosiypb.FoodDetailRequest{
		Id:      id,
		Type:    c.Query("type"),
		Portion: int32(portion),
	}

	res, err := h.service.GetFoodDetail(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ingredients := make([]mailmodels.Ingredient, 0)

	for _, ing := range res.Ingredients {

		ingredients = append(ingredients, mailmodels.Ingredient{
			Name:   ing.Name,
			Amount: ing.Amount,
		})
	}

	food := mailmodels.FoodDetail{
		ID:                  res.Id,
		Type:                res.Type,
		RestaurantID:        res.RestaurantId,
		Name:                res.Name,
		Description:         res.Description,
		ImageURL:            res.ImageUrl,
		VideoURL:            res.VideoUrl,
		Country:             res.Country,
		MealTime:            res.MealTime,
		Kcal:                res.Kcal,
		Protein:             res.Protein,
		Fat:                 res.Fat,
		Carbs:               res.Carbs,
		Portion:             res.Portion,
		TotalKcal:           float64(res.TotalKcal),
		CookingTimeMinutes:  res.CookingTimeMinutes,
		Ingredients:         ingredients,
		Steps:               res.Steps,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    food,
	})
}