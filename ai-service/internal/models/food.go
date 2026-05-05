package models

type MealRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Country     string  `json:"country"`
    MealTime    string  `json:"meal_time"`
    Kcal        float32 `json:"kcal"`
    Protein     float32 `json:"protein"`
    Fat         float32 `json:"fat"`
    Carbs       float32 `json:"carbs"`
    Portion     int32   `json:"portion"`
}

type MealResponse struct {
    Portion            int32        `json:"portion"`
    TotalKcal          float32      `json:"total_kcal"`
    CookingTimeMinutes int32        `json:"cooking_time_minutes"`
    Ingredients        []Ingredient `json:"ingredients"`
    Steps              []string     `json:"steps"`
}

type Ingredient struct {
    Name   string `json:"name"`
    Amount string `json:"amount"`
}