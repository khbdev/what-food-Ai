package models

type CityStat struct {
	City  string
	Count int64
}

type DashboardStats struct {
	TotalUsers   int64
	ActiveUsers  int64
	CityStats    []CityStat
}

