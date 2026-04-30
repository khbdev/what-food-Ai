package models

import "time"

type Restaurant struct {
	ID             int64     `json:"id"`
	RestaurantName string    `json:"restaurant_name"`
	Description    string    `json:"description"`
	ImageURL       string    `json:"image_url"`
	CreatedAt      time.Time `json:"created_at"`
}a