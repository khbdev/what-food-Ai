package models

import "time"

type IngredientP struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Name       string    `json:"name"`
	Quantity   float64   `json:"quantity"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}