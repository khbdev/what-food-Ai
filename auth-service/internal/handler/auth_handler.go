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

