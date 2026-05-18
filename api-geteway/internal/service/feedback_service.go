package service

import (
	"context"
	"errors"

	"api-geteway/internal/client"
	pb "github.com/khbdev/what-food-proto/proto/feedback"
)

type FeedbackService struct {
	feedbackClient *client.FeedbackClient
}

func NewFeedbackService(c *client.FeedbackClient) *FeedbackService {
	return &FeedbackService{
		feedbackClient: c,
	}
}

func (s *FeedbackService) AnalyzeNutrition(
	ctx context.Context,
	req *pb.NutritionRequest,
) (*pb.NutritionResponse, error) {

	// 🔥 validation
	if req == nil {
		return nil, errors.New("request is nil")
	}

	if len(req.Days) == 0 {
		return nil, errors.New("days is required")
	}

	// optional: extra validation
	for _, d := range req.Days {
		if d.Kcal < 0 || d.Fat < 0 || d.Carbs < 0 || d.Protein < 0 {
			return nil, errors.New("invalid nutrition values")
		}
	}

	// gRPC call
	return s.feedbackClient.AnalyzeNutrition(req)
}