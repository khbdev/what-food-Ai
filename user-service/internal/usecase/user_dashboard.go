package usecase

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/models"
)

type userDashboardUsecase struct {
	repo domain.UserDashboard
}

func NewUserDashboardUsecase(repo domain.UserDashboard) domain.UserDashboard {
	return &userDashboardUsecase{
		repo: repo,
	}
}

func (u *userDashboardUsecase) GetDashboardStats(ctx context.Context,) (*models.DashboardStats, error) {
	return u.repo.GetDashboardStats(ctx)
}