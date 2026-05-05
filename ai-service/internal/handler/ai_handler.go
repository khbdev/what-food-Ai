package handler

import (
	"ai-service/internal/usecase"
	pb ""
	"context"
)

type AIHandler struct {
	pb.UnimplementedAiServiceServer
	usecase *usecase.AIUsecase
}

func NewAIHandler(usecase *usecase.AIUsecase) *AIHandler {
	return &AIHandler{usecase: usecase}
}

func (h *AIHandler) AnalyzeMeal(ctx context.Context, req *pb.MealRequest) (*pb.MealResponse, error) {
	res, err := h.usecase.AnalyzeMeal(
		ctx,
		req.Name,
		req.Description,
		req.Country,
		req.MealTime,
		req.Kcal,
		req.Protein,
		req.Fat,
		req.Carbs,
		req.Portion,
	)
	if err != nil {
		return nil, err
	}

	var ingredients []*pb.Ingredient
	for _, ing := range res.Ingredients {
		ingredients = append(ingredients, &pb.Ingredient{
			Name:   ing.Name,
			Amount: ing.Amount,
		})
	}

	return &pb.MealResponse{
		Portion:            res.Portion,
		TotalKcal:          res.TotalKcal,
		CookingTimeMinutes: res.CookingTimeMinutes,
		Ingredients:        ingredients,
		Steps:              res.Steps,
	}, nil
}

func (h *AIHandler) AnalyzeNutrition(ctx context.Context, req *pb.NutritionRequest) (*pb.NutritionResponse, error) {
	res, err := h.usecase.AnalyzeNutrition(
		ctx,
		req.Period,
		req.AvgKcal,
		req.AvgProtein,
		req.AvgFat,
		req.AvgCarbs,
	)
	if err != nil {
		return nil, err
	}

	return &pb.NutritionResponse{
		Feedback: res.Feedback,
		Level:    res.Level,
	}, nil
}