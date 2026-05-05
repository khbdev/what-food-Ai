package main

import (
	aimodel "ai-service/internal/ai-model"
	"ai-service/internal/models"
	"ai-service/internal/usecase"
	"ai-service/pkg/env"
	"context"
	"fmt"
	"log"
	"time"
)


func main() {
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()

	aiUse := usecase.NewAIUsecase()
}