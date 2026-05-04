package models

// ======================
// COMMON
// ======================

type Empty struct{}

type GetByIDRequest struct {
	ID int64 `json:"id"`
}

// ======================
// RECIPE
// ======================

type Recipe struct {
	ID           int64   `json:"id"`
	RestaurantID int64   `json:"restaurant_id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`

	Country  string `json:"country"`
	MealTime string `json:"meal_time"`

	Kcal    int32   `json:"kcal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
}

type CreateRecipeRequest struct {
	Recipe Recipe `json:"recipe"`
}

type UpdateRecipeRequest struct {
	Recipe Recipe `json:"recipe"`
}

type RecipeResponse struct {
	Recipe Recipe `json:"recipe"`
}

type RecipeListResponse struct {
	Recipes []Recipe `json:"recipes"`
}

// ======================
// SALAD
// ======================

type Salad struct {
	ID           int64 `json:"id"`
	RestaurantID int64 `json:"restaurant_id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`

	Country  string `json:"country"`
	MealTime string `json:"meal_time"`

	Kcal    int32   `json:"kcal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
}

type CreateSaladRequest struct {
	Salad Salad `json:"salad"`
}

type UpdateSaladRequest struct {
	Salad Salad `json:"salad"`
}

type SaladResponse struct {
	Salad Salad `json:"salad"`
}

type SaladListResponse struct {
	Salads []Salad `json:"salads"`
}

// ======================
// RESTAURANT
// ======================

type Restaurant struct {
	ID             int64  `json:"id"`
	RestaurantName string `json:"restaurant_name"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
}

type CreateRestaurantRequest struct {
	RestaurantName string `json:"restaurant_name"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
}

type CreateRestaurantResponse struct {
	ID int64 `json:"id"`
}

type UpdateRestaurantRequest struct {
	RestaurantName string `json:"restaurant_name" binding:"required"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
}
type RestaurantResponse struct {
	Restaurant Restaurant `json:"restaurant"`
}

type RestaurantListResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
}

// ======================
// FILTER
// ======================

type FoodFilterRequest struct {
	Country       string `json:"country"`
	MealTime      string `json:"meal_time"`
	MaxKcal       int32  `json:"max_kcal"`
	IncludeSalads bool   `json:"include_salads"`
}

type FoodItem struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"` // recipe | salad

	RestaurantID int64 `json:"restaurant_id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`

	Country  string `json:"country"`
	MealTime string `json:"meal_time"`

	Kcal    int32   `json:"kcal"`
	Protein float64 `json:"protein"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
}

type FoodListResponse struct {
	Items []FoodItem `json:"items"`
}