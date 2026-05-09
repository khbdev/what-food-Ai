package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"github.com/khbdev/what-food-proto/proto/statik
	"google.golang.org/grpc/credentials/insecure"

	nutritionpb "github.com/khbdev/what-food-proto/proto/statik"
)

type NutritionClient struct {
	conn *grpc.ClientConn
	svc  nutritionpb.NutritionServiceClient
}

// =========================
// CONFIG
// =========================

const (
	timeoutNutrition = 5 * time.Second
)

// =========================
// INIT (FAIL FAST)
// =========================

func NewNutritionClient(addr string) (*NutritionClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutNutrition)
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

	log.Println("✅ Nutrition service connected:", addr)

	return &NutritionClient{
		conn: conn,
		svc:  nutritionpb.NewNutritionServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *NutritionClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *NutritionClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutNutrition)
}

// =========================
// METHODS
// =========================

func (c *NutritionClient) GetWeeklyNutrition(
	req *nutritionpb.WeeklyNutritionRequest,
) (*nutritionpb.WeeklyNutritionResponse, error) {

	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetWeeklyNutrition(ctx, req)
}