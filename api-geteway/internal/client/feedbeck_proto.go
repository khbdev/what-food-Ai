package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/khbdev/what-food-proto/proto/feedback"
)

type FeedbackClient struct {
	conn *grpc.ClientConn
	svc  pb.MailServiceClient
}

const timeoutFeedback = 5 * time.Second

// =========================
// INIT
// =========================

func NewFeedbackClient(addr string) (*FeedbackClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutFeedback)
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

	log.Println("✅ Feedback service connected:", addr)

	return &FeedbackClient{
		conn: conn,
		svc:  pb.NewMailServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *FeedbackClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *FeedbackClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutFeedback)
}

// =========================
// ANALYZE FEEDBACK
// =========================

func (c *FeedbackClient) AnalyzeNutrition(
	req *pb.NutritionRequest,
) (*pb.NutritionResponse, error) {

	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.AnalyzeNutrition(ctx, req)
}