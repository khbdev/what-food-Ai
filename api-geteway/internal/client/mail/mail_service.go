package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	asosiypb "github.com/khbdev/what-food-proto/proto/asosiy"
)

type FoodClient struct {
	conn *grpc.ClientConn
	svc  asosiypb.FoodServiceClient
}

// =========================
// CONFIG
// =========================

const (
	timeoutFood = 5 * time.Second
)

// =========================
// INIT (FAIL FAST)
// =========================

func NewFoodClient(addr string) (*FoodClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutFood)
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

	log.Println("✅ Food service connected:", addr)

	return &FoodClient{
		conn: conn,
		svc:  asosiypb.NewFoodServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *FoodClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *FoodClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutFood)
}

// =========================
// METHODS
// =========================

// Filter foods
func (c *FoodClient) FilterFood(req *asosiypb.FoodFilterRequest) (*asosiypb.FoodListResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.FilterFood(ctx, req)
}

// Get food detail
func (c *FoodClient) GetFoodDetail(req *asosiypb.FoodDetailRequest) (*asosiypb.FoodDetailResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetFoodDetail(ctx, req)
}