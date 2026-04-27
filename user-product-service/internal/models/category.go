package models

import "time"

type Category struct {
    ID        int64     `db:"id"`
    Name      string    `db:"name"`
    CreatedAt time.Time `db:"created_at"`
}


type Ingredient struct {
	ID       int64
	Name     string
	Quantity float64
}

type CategoryWithIngredients struct {
	CategoryID int64
	Name       string
	Items      []Ingredient
}