package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "github.com/khbdev/what-food-proto/proto/auth"
)

type AuthClient struct {
	conn *grpc.ClientConn
	svc  authpb.AuthServiceClient
}

const timeout = 5 * time.Second

// =========================
// INIT (ONLY CONNECT)
// =========================

func NewAuthClient(addr string) (*AuthClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
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

	log.Println("✅ Auth service connected:", addr)

	return &AuthClient{
		conn: conn,
		svc:  authpb.NewAuthServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *AuthClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// =========================
// METHODS
// =========================

func (c *AuthClient) Register(req *authpb.RegisterRequest) (*authpb.SimpleResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.Register(ctx, req)
}

func (c *AuthClient) Login(req *authpb.LoginRequest) (*authpb.SimpleResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.Login(ctx, req)
}

func (c *AuthClient) VerifyOTP(req *authpb.VerifyRequest) (*authpb.AuthResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.VerifyOTP(ctx, req)
}