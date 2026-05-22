package models

// =======================
// DATA MODEL
// =======================

type DayNutrition struct {
	Day     string  `json:"day"`
	Kcal    float64 `json:"kcal"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Protein float64 `json:"protein"`
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