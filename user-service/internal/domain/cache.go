package domain

import (
	"context"
	"time"
	"user-service/internal/models"
)



type UserCache interface {
	// Read-through uchun
	GetUser(ctx context.Context, id uint) (*models.User, error)

	// Write-through uchun
	SetUser(ctx context.Context, user *models.User, ttl time.Duration) error

	// Delete cache
	DeleteUser(ctx context.Context, id uint) error
}