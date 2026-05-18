package models


type DayDTO struct {
	Day     string
	Kcal    float32
	Fat     float32
	Carbs   float32
	Protein float32
}

type NutritionRequestDTO struct {
	Period string
	Days   []DayDTO
}
