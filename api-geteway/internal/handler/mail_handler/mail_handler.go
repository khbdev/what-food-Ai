package mailhandler

import (
	"net/http"
	"strconv"

	"api-geteway/internal/models"
	"api-geteway/internal/service/mail"

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
// GET /foods/filter?country=USA&meal_time=lunch&has_salad=true&kcal_limit=500
// =========================

func (h *FoodHandler) FilterFood(c *gin.Context) {

	// query parse
	kcalLimit, _ := strconv.Atoi(c.DefaultQuery("kcal_limit", "0"))

	hasSalad, _ := strconv.ParseBool(c.DefaultQuery("has_salad", "false"))

	req := &models.FoodFilter{
		Country:   c.Query("country"),
		MealTime:  c.Query("meal_time"),
		HasSalad:  hasSalad,
		KcalLimit: int32(kcalLimit),
	}

	// grpc request
	grpcReq := &asosiypb.FoodFilterRequest{
		Country:   req.Country,
		MealTime:  req.MealTime,
		HasSalad:  req.HasSalad,
		KcalLimit: req.KcalLimit,
	}

	// service call
	res, err := h.service.FilterFood(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// mapping response
	foods := make([]models.FoodItem, 0)

	for _, item := range res.Items {

		foods = append(foods, models.FoodItem{
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
// GET /foods/detail?id=1&type=ai&portion=2
// =========================

func (h *FoodHandler) GetFoodDetail(c *gin.Context) {

	// parse params
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

	req := &models.FoodDetailRequest{
		ID:      id,
		Type:    c.Query("type"),
		Portion: int32(portion),
	}

	// grpc request
	grpcReq := &asosiypb.FoodDetailRequest{
		Id:      req.ID,
		Type:    req.Type,
		Portion: req.Portion,
	}

	// service call
	res, err := h.service.GetFoodDetail(c.Request.Context(), grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ingredients mapping
	ingredients := make([]models.Ingredient, 0)

	for _, ing := range res.Ingredients {
		ingredients = append(ingredients, models.Ingredient{
			Name:   ing.Name,
			Amount: ing.Amount,
		})
	}

	// final response
	food := models.FoodDetail{
		ID:                   res.Id,
		Type:                 res.Type,
		RestaurantID:         res.RestaurantId,
		Name:                 res.Name,
		Description:          res.Description,
		ImageURL:             res.ImageUrl,
		VideoURL:             res.VideoUrl,
		Country:              res.Country,
		MealTime:             res.MealTime,
		Kcal:                 res.Kcal,
		Protein:              res.Protein,
		Fat:                  res.Fat,
		Carbs:                res.Carbs,
		Portion:              res.Portion,
		TotalKcal:            float64(res.TotalKcal),
		CookingTimeMinutes:   res.CookingTimeMinutes,
		Ingredients:          ingredients,
		Steps:                res.Steps,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    food,
	})
}