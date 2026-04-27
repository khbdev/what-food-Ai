package models

import "time"

type Ingredient struct {
    ID         int64     `db:"id"`
    UserID     int64     `db:"user_id"`
    Name       string    `db:"name"`
    Quantity   float64   `db:"quantity"`
    CategoryID int64     `db:"category_id"`
    CreatedAt  time.Time `db:"created_at"`
}