package models

// =========================
// REQUEST
// =========================

type WeeklyNutritionRequest struct {
	UserID int64 `json:"user_id"`
}

// =========================
// RESPONSE ITEM
// =========================

type DailyNutrition struct {
	Day     string  `json:"day"`
	Kcal    float64 `json:"kcal"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Protein float64 `json:"protein"`
}

// =========================
// RESPONSE
// =========================

type WeeklyNutritionResponse struct {
	Days []DailyNutrition `json:"days"`
}