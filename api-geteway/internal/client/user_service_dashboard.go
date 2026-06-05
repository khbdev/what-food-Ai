package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dashboardpb "github.com/khbdev/what-food-proto/proto/dashboard"
)

type DashboardClient struct {
	conn *grpc.ClientConn
	svc  dashboardpb.UserDashboardServiceClient
}

// =========================
// CONFIG
// =========================

const (
	timeoutDASHBOARD = 5 * time.Second
)

// =========================
// INIT (FAIL FAST)
// =========================

func NewDashboardClient(addr string) (*DashboardClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDASHBOARD)
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

	log.Println("✅ Dashboard service connected:", addr)

	return &DashboardClient{
		conn: conn,
		svc:  dashboardpb.NewUserDashboardServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *DashboardClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *DashboardClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutDASHBOARD)
}

// =========================
// METHODS
// =========================

func (c *DashboardClient) GetDashboardStats() (*dashboardpb.DashboardStatsResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetDashboardStats(ctx, &dashboardpb.Empty{})
}