type NutritionRequest struct {
    Period     string  `json:"period"`
    AvgKcal    float32 `json:"avg_kcal"`
    AvgProtein float32 `json:"avg_protein"`
    AvgFat     float32 `json:"avg_fat"`
    AvgCarbs   float32 `json:"avg_carbs"`
}

type NutritionResponse struct {
    Feedback string `json:"feedback"`
    Level    string `json:"level"` // "danger", "bad", "normal", "good"
}