package service

import (
	"api-geteway/internal/client"
	"api-geteway/internal/models"
	"context"
)

type DashboardService struct {
	dashboardClient *client.DashboardClient
}

func NewDashboardService(c *client.DashboardClient) *DashboardService {
	return &DashboardService{
		dashboardClient: c,
	}
}


func (s *DashboardService) GetDashboardStats(ctx context.Context) (*models.DashboardStats, error) {

	res, err := s.dashboardClient.GetDashboardStats()
	if err != nil {
		return nil, err
	}

	// mapping proto -> internal model
	stats := &models.DashboardStats{
		TotalUsers:  res.TotalUsers,
		ActiveUsers: res.ActiveUsers,
		CityStats:   make([]models.CityStat, 0, len(res.CityStats)),
	}

	for _, c := range res.CityStats {
		stats.CityStats = append(stats.CityStats, models.CityStat{
			City:  c.City,
			Count: c.Count,
		})
	}

	return stats, nil
}