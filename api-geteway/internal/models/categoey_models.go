package models

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Ingredient struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

type CategoryWithIngredients struct {
	CategoryID int64        `json:"category_id"`
	Name       string       `json:"name"`
	Items      []Ingredient `json:"items"`
}