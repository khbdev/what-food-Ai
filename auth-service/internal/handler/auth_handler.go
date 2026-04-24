package handler

import (
	"context"

	"auth-service/internal/domain"
	

	authpb  "github.com/khbdev/what-food-proto/proto/auth"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: u,
	}
}

func (h *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.SimpleResponse, error) {

	err := h.authUsecase.Register(models.RegisterRequest{
		FullName: req.FullName,
		Phone:    req.Phone,
		Age:      int(req.Age),
		Address:  req.Address,
	})

	if err != nil {
		return nil, err
	}

	return &authpb.SimpleResponse{
		Message: "success",
	}, nil
}