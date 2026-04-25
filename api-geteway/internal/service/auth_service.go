package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	authpb "github.com/khbdev/what-food-proto/proto/authpb"
)

type AuthService struct {
	authClient *client.AuthClient
}

func NewAuthService(c *client.AuthClient) *AuthService {
	return &AuthService{authClient: c}
}

// =========================
// REGISTER
// =========================

func (s *AuthService) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.SimpleResponse, error) {

	// 🔥 validation
	if strings.TrimSpace(req.FullName) == "" {
		return nil, errors.New("full_name is required")
	}

	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("phone is required")
	}

	if len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	if req.Age <= 0 || req.Age > 120 {
		return nil, errors.New("invalid age")
	}

	if strings.TrimSpace(req.Address) == "" {
		return nil, errors.New("address is required")
	}

	// gRPC call
	return s.authClient.Register(req)
}

// =========================
// LOGIN
// =========================

func (s *AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.SimpleResponse, error) {

	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("phone is required")
	}

	if len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	return s.authClient.Login(req)
}

// =========================
// VERIFY OTP
// =========================

func (s *AuthService) VerifyOTP(ctx context.Context, req *authpb.VerifyRequest) (*authpb.AuthResponse, error) {

	if req.Otp <= 0 {
		return nil, errors.New("otp is required")
	}

	if req.Otp < 100000 || req.Otp > 999999 {
		return nil, errors.New("otp must be 6 digits")
	}

	return s.authClient.VerifyOTP(req)
}