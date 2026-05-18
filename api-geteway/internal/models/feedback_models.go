package models

// =======================
// DATA MODEL
// =======================

type DayNutrition struct {
	Day     string  `json:"day"`
	Kcal    int32   `json:"kcal"`
	Fat     int32   `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Protein int32   `json:"protein"`
}

// =======================
// REQUEST (API Gateway -> Mail Service)
// =======================

type NutritionRequest struct {
	Days []DayNutrition `json:"days"`
}

// =======================
// RESPONSE (Mail Service -> API Gateway)
// =======================

type NutritionResponse struct {
	Feedback string `json:"feedback"`
	Level    string `json:"level"`
}