package domain

import (
	"auth-service/internal/models"
)

type AuthUsecase interface {
	Register(req models.RegisterRequest) error
	Login(req models.LoginRequest) error
	Verify(code int64) (string, string, error)
}