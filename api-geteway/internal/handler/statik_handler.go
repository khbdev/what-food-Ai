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

