package repository

import (
	"context"
	"time"
	"user-service/internal/domain"
	"user-service/internal/models"

	"gorm.io/gorm"
)





type userDashboard struct {
	db *gorm.DB
}

func NewUserDashboard(db *gorm.DB) domain.UserDashboard {
	return &userDashboard{
		db: db,
	}
}


func (r *userDashboard) GetDashboardStats(ctx context.Context) (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	// 1. Total users
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Count(&stats.TotalUsers).Error; err != nil {
		return nil, err
	}

	// 2. Active users (oxirgi 10 daqiqada aktiv bo‘lgan deb olamiz)
	activeThreshold := time.Now().Add(-10 * time.Minute)

	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("updated_at >= ?", activeThreshold).
		Count(&stats.ActiveUsers).Error; err != nil {
		return nil, err
	}

	// 3. City bo‘yicha grouping
	type rawCity struct {
		Address string
		Count   int64
	}

	var raw []rawCity

	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("address, COUNT(*) as count").
		Group("address").
		Order("count DESC").
		Find(&raw).Error; err != nil {
		return nil, err
	}

	// map → clean DTO
	stats.CityStats = make([]models.CityStat, 0, len(raw))

	for _, r := range raw {
		stats.CityStats = append(stats.CityStats, models.CityStat{
			City:  r.Address,
			Count: r.Count,
		})
	}

	return stats, nil
}