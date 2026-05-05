package aimodel

import (
	"ai-service/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type GroqClient struct {
	client openai.Client
}

func NewGroqClient() *GroqClient {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("GROQ_API_KEY")),
		option.WithBaseURL("https://api.groq.com/openai/v1"),
	)
	return &GroqClient{client: client}
}

func (g *GroqClient) AnalyzeMeal(ctx context.Context, req models.MealRequest) (*models.MealResponse, error) {
	prompt := fmt.Sprintf(`You are a cooking assistant.
Return ONLY a raw JSON object with these fields:
- portion: %d
- total_kcal: %.0f
- cooking_time_minutes: integer
- ingredients: array of {name, amount}
- steps: array of strings

Meal: %s, %s, %s, %s
Nutrition per portion: %.0f kcal, %.0fg protein, %.0fg fat, %.0fg carbs`,
		req.Portion,
		req.Kcal*float32(req.Portion),
		req.Name, req.Description, req.Country, req.MealTime,
		req.Kcal, req.Protein, req.Fat, req.Carbs)

	resp, err := g.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: "llama-3.3-70b-versatile",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})
	if err != nil {
		return nil, err
	}

	raw := resp.Choices[0].Message.Content
	log.Println("RAW:", raw)

	var result models.MealResponse
	if err := json.Unmarshal([]byte(raw), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (g *GroqClient) AnalyzeNutrition(ctx context.Context, req models.NutritionRequest) (*models.NutritionResponse, error) {
	prompt := fmt.Sprintf(`You are a nutrition analyst.
Respond ONLY in this exact JSON format, no extra text, no markdown, no backticks:

{
  "feedback": "...",
  "level": "danger/bad/normal/good"
}

Period: %s
Avg kcal: %.0f
Avg protein: %.0fg
Avg fat: %.0fg
Avg carbs: %.0fg`,
		req.Period, req.AvgKcal, req.AvgProtein, req.AvgFat, req.AvgCarbs)

	resp, err := g.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: "llama-3.3-70b-versatile",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})
	if err != nil {
		return nil, err
	}

	var result models.NutritionResponse
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		return nil, err
	}

	return &result, nil
}