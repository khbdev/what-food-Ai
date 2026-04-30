package models

import "time"

type Recipe struct {
	ID          int64     `json:"id"`
	resipe_id int64 
	Name        string    `json:"name"`
	Description string    `json:"description"`

	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`

	Country  string `json:"country"`
	MealTime string `json:"meal_time"`

	IsSalad bool `json:"is_salad"`

	Kcal    int     `json:"kcal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`

	CreatedAt time.Time `json:"created_at"`
}