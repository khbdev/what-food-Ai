package domain

import (
	"context"
	"user-service/internal/models"
)




type UserDashboard interface {
	GetDashboardStats(ctx context.Context) (*models.DashboardStats, error)
}