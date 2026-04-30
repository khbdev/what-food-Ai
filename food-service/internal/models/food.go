package models

import "time"

type Recipe struct {
	ID          int64  `json:"id"`
	RestaurantID int64  `json:"restaurant_id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`

	Country  string `json:"country"`
	MealTime string `json:"meal_time"`



	Kcal    int     `json:"kcal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`

	CreatedAt time.Time `json:"created_at"`
}