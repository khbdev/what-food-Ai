package handler

import (
	"context"

	"auth-service/internal/domain"
	"auth-service/internal/models"

	authpb "github.com/khbdev/what-food-proto/proto/authpb"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: u,
	}
}