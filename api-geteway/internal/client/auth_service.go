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

// =====================
// CONFIG
// =====================

const (
	timeout = 5 * time.Second
)

// =====================
// INIT
// =====================

func NewAuthClient(addr string) (*AuthClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // 🔥 muhim: real connect kutadi
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

// =====================
// CLOSE
// =====================

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

// =========================
// RETRY WRAPPER
// =========================

func withRetry(fn func() error) error {
	var err error

	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		time.Sleep(retryDelay)
	}

	return err
}

// =========================
// METHODS
// =========================

func (c *AuthClient) Register(req *authpb.RegisterRequest) (*authpb.SimpleResponse, error) {
	var (
		res *authpb.SimpleResponse
		err error
	)

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.Register(ctx, req)
		return err
	})

	return res, err
}

func (c *AuthClient) Login(req *authpb.LoginRequest) (*authpb.SimpleResponse, error) {
	var (
		res *authpb.SimpleResponse
		err error
	)

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.Login(ctx, req)
		return err
	})

	return res, err
}

func (c *AuthClient) VerifyOTP(req *authpb.VerifyRequest) (*authpb.AuthResponse, error) {
	var (
		res *authpb.AuthResponse
		err error
	)

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.VerifyOTP(ctx, req)
		return err
	})

	return res, err
}