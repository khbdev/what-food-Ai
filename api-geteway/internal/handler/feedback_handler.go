package handler

import (
	"net/http"

	"api-geteway/internal/models"
	"api-geteway/internal/service"

	pb "github.com/khbdev/what-food-proto/proto/feedback"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	service *service.FeedbackService
}

func NewFeedbackHandler(s *service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		service: s,
	}
}


func (h *FeedbackHandler) AnalyzeNutrition(c *gin.Context) {

	var req models.NutritionRequest

	// 1. bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 2. convert models → proto
	pbReq := &pb.NutritionRequest{
		Days: make([]*pb.DayNutrition, 0, len(req.Days)),
	}

	for _, d := range req.Days {
		pbReq.Days = append(pbReq.Days, &pb.DayNutrition{
			Day:     d.Day,
			Kcal:    d.Kcal,
			Fat:     d.Fat,
			Carbs:   d.Carbs,
			Protein: d.Protein,
		})
	}

	// 3. call service (business layer)
	resp, err := h.service.AnalyzeNutrition(c.Request.Context(), pbReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 4. response → models
	result := models.NutritionResponse{
		Feedback: resp.Feedback,
		Level:    resp.Level,
	}

	c.JSON(http.StatusOK, result)
}