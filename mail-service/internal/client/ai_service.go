package client

import (
	"context"
	"log"
	"time"

	aipb "github.com/khbdev/what-food-proto/proto/food"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const aiTimeout = 5 * time.Second

type AiClient struct {
	conn *grpc.ClientConn

	Meal      aipb.AiServiceClient
	Nutrition aipb.AiServiceClient
}

// =========================
// INIT
// =========================

func NewAiClient(addr string) (*AiClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), aiTimeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, err
	}

	log.Println("✅ AI service connected:", addr)

	client := aipb.NewAiServiceClient(conn)

	return &AiClient{
		conn:      conn,
		Meal:      client,
		Nutrition: client,
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *AiClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT
// =========================

func (c *AiClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), aiTimeout)
}

// =========================
// ANALYZE MEAL
// =========================

func (c *AiClient) AnalyzeMeal(
	req *aipb.MealRequest,
) (*aipb.MealResponse, error) {

	ctx, cancel := c.ctx()
	defer cancel()

	return c.Meal.AnalyzeMeal(ctx, req)
}

// =========================
// ANALYZE NUTRITION
// =========================

func (c *AiClient) AnalyzeNutrition(
	req *aipb.NutritionRequest,
) (*aipb.NutritionResponse, error) {

	ctx, cancel := c.ctx()
	defer cancel()

	return c.Nutrition.AnalyzeNutrition(ctx, req)
}