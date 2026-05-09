package handler

import (
	"net/http"

	"api-geteway/internal/models"
	"api-geteway/internal/service"

	nutritionpb "github.com/khbdev/what-food-proto/proto/statik"
	"github.com/gin-gonic/gin"
)

type NutritionHandler struct {
	service *service.NutritionService
}

// =========================
// INIT
// =========================

func NewNutritionHandler(s *service.NutritionService) *NutritionHandler {
	return &NutritionHandler{
		service: s,
	}
}

// =========================
// GET WEEKLY NUTRITION
// =========================

func (h *NutritionHandler) GetWeeklyNutrition(c *gin.Context) {

	var req models.WeeklyNutritionRequest

	// =========================
	// BIND JSON
	// =========================

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// CALL SERVICE
	// =========================

	res, err := h.service.GetWeeklyNutrition(
		c.Request.Context(),
		&nutritionpb.WeeklyNutritionRequest{
			UserId: req.UserID,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// MAPPING RESPONSE
	// =========================

	var days []models.DailyNutrition

	for _, d := range res.Days {
		days = append(days, models.DailyNutrition{
			Day:     d.Day,
			Kcal:    d.Kcal,
			Fat:     d.Fat,
			Carbs:   d.Carbs,
			Protein: d.Protein,
		})
	}

	// =========================
	// RESPONSE
	// =========================

	c.JSON(http.StatusOK, models.WeeklyNutritionResponse{
		Days: days,
	})
}