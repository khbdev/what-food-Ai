package models

type RecipeFilter struct {
	Country       string `json:"country"`
	MealTime      string `json:"meal_time"`
	MaxKcal       int    `json:"max_kcal"`
	IncludeSalads bool   `json:"include_salads"`
}