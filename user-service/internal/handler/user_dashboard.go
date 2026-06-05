package handler

import (
	"context"
	"user-service/internal/domain"

	dashboardpb "github.com/khbdev/what-food-proto/proto/dashboard"
)

type DashboardHandler struct {
	dashboardpb.UnimplementedUserDashboardServiceServer
	usecase domain.UserDashboard
}

func NewDashboardHandler(usecase domain.UserDashboard) *DashboardHandler {
	return &DashboardHandler{
		usecase: usecase,
	}
}


func (h *DashboardHandler) GetDashboardStats(
	ctx context.Context,
	req *dashboardpb.Empty,
) (*dashboardpb.DashboardStatsResponse, error) {

	stats, err := h.usecase.GetDashboardStats(ctx)
	if err != nil {
		return nil, err
	}

	// map: models → proto
	res := &dashboardpb.DashboardStatsResponse{
		TotalUsers:  stats.TotalUsers,
		ActiveUsers: stats.ActiveUsers,
		CityStats:   make([]*dashboardpb.CityStat, 0, len(stats.CityStats)),
	}

	for _, c := range stats.CityStats {
		res.CityStats = append(res.CityStats, &dashboardpb.CityStat{
			City:  c.City,
			Count: c.Count,
		})
	}

	return res, nil
}