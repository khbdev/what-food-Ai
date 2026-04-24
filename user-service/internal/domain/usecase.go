package domain

import (
	"context"
	"user-service/internal/models"
)


type UserUsecase interface {
	Create(ctx context.Context, user *models.User) (*models.User, error) // ← *models.User qaytaradi
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
	GetAll(ctx context.Context, req *models.GetAllUsersRequest) (*models.GetAllUsersResponse, error)
	Update(ctx context.Context, id uint, req *models.User) error
	Delete(ctx context.Context, id uint) error
}