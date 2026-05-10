package mailhandler

import (
	mailmodels "api-geteway/internal/models/mail_models"
	"api-geteway/internal/service/mail"
	"net/http"

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

	// =========================
	// PARSE JSON
	// =========================

	var body mailmodels.FoodFilter

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// DOMAIN -> PROTO
	// =========================

	req := &asosiypb.FoodFilterRequest{
		Country:   body.Country,
		MealTime:  body.MealTime,
		HasSalad:  body.HasSalad,
		KcalLimit: body.KcalLimit,
	}

	// =========================
	// SERVICE CALL
	// =========================

	res, err := h.service.FilterFood(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// RESPONSE MAPPING
	// =========================

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

	// =========================
	// PARSE JSON
	// =========================

	var body mailmodels.FoodDetailRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// DOMAIN -> PROTO
	// =========================

	req := &asos

	// =========================
	// SERVICE CALL
	// =========================

	res, err := h.service.GetFoodDetail(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// INGREDIENTS MAPPING
	// =========================

	ingredients := make([]mailmodels.Ingredient, 0)

	for _, ing := range res.Ingredients {

		ingredients = append(ingredients, mailmodels.Ingredient{
			Name:   ing.Name,
			Amount: ing.Amount,
		})
	}

	// =========================
	// RESPONSE MAPPING
	// =========================

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