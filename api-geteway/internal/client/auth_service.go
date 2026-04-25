package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "github.com/khbdev/what-food-proto/proto/auth"
)

type AuthClient struct {
	conn *grpc.ClientConn
	svc  authpb.AuthServiceClient
}

// =========================
// INIT / CLOSE
// =========================

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &AuthClient{
		conn: conn,
		svc:  authpb.NewAuthServiceClient(conn),
	}, nil
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONFIG
// =========================

const (
	requestTimeout = 5 * time.Second
	maxRetries     = 3
	retryDelay     = 200 * time.Millisecond
)

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
	var res *authpb.SimpleResponse
	var err error

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.Register(ctx, req)
		return err
	})

	return res, err
}

func (c *AuthClient) Login(req *authpb.LoginRequest) (*authpb.SimpleResponse, error) {
	var res *authpb.SimpleResponse
	var err error

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.Login(ctx, req)
		return err
	})

	return res, err
}

func (c *AuthClient) VerifyOTP(req *authpb.VerifyRequest) (*authpb.AuthResponse, error) {
	var res *authpb.AuthResponse
	var err error

	err = withRetry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		res, err = c.svc.VerifyOTP(ctx, req)
		return err
	})

	return res, err
}