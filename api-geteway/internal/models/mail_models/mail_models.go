package mailmodels

// =========================
// FOOD ITEM (LIST)
// =========================

type FoodItem struct {
	ID           int64   `json:"id"`
	Type         string  `json:"type"`
	RestaurantID int64   `json:"restaurant_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	VideoURL     string  `json:"video_url"`
	Country      string  `json:"country"`
	MealTime     string  `json:"meal_time"`
	Kcal         int32   `json:"kcal"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbs        float64 `json:"carbs"`
}

// =========================
// FILTER REQUEST (DOMAIN)
// =========================

type FoodFilter struct {
	Country   string `json:"country"`
	MealTime  string `json:"meal_time"`
	HasSalad  bool   `json:"has_salad"`
	KcalLimit int32  `json:"kcal_limit"`
}
// =========================
// DETAIL REQUEST
// =========================

type FoodDetailRequest struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"`
	Portion  int32  `json:"portion"`
	UserID   int64  `json:"user_id"`
}
// =========================
// INGREDIENT
// =========================

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

// =========================
// FOOD DETAIL (FULL)
// =========================

type FoodDetail struct {
	// base food data
	ID           int64   `json:"id"`
	Type         string  `json:"type"`
	RestaurantID int64   `json:"restaurant_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	VideoURL     string  `json:"video_url"`
	Country      string  `json:"country"`
	MealTime     string  `json:"meal_time"`
	Kcal         int32   `json:"kcal"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbs        float64 `json:"carbs"`

	// AI calculated fields
	Portion             int32        `json:"portion"`
	TotalKcal           float64      `json:"total_kcal"`
	CookingTimeMinutes  int32        `json:"cooking_time_minutes"`
	Ingredients         []Ingredient `json:"ingredients"`
	Steps               []string     `json:"steps"`
}