package usecase


func (u *FoodUsecase) GetFoodDetailAndAnalyze(
	id int64,
	foodType string,
	portion int32,
) (*aipb.MealResponse, error) {

	// =========================
	// VALIDATION
	// =========================

	if id == 0 {
		return nil, errors.New("id is required")
	}

	if foodType != "recipe" && foodType != "salad" {
		return nil, errors.New("invalid type")
	}

	// =========================
	// FETCH FOOD
	// =========================

	var (
		name, desc, country, mealTime string
		kcal, protein, fat, carbs      float64
	)

	if foodType == "recipe" {

		r, err := u.foodClient.GetRecipeByID(id)
		if err != nil {
			return nil, err
		}

		name = r.Name
		desc = r.Description
		country = r.Country
		mealTime = r.MealTime
		kcal = float64(r.Kcal)
		protein = r.Protein
		fat = r.Fat
		carbs = r.Carbs
	}

	if foodType == "salad" {

		s, err := u.foodClient.GetSaladByID(id)
		if err != nil {
			return nil, err
		}

		name = s.Name
		desc = s.Description
		country = s.Country
		mealTime = s.MealTime
		kcal = float64(s.Kcal)
		protein = s.Protein
		fat = s.Fat
		carbs = s.Carbs
	}

	// =========================
	// AI REQUEST
	// =========================

	aiReq := &aipb.MealRequest{
		Name:        name,
		Description: desc,
		Country:     country,
		MealTime:    mealTime,
		Kcal:        float32(kcal),
		Protein:     float32(protein),
		Fat:         float32(fat),
		Carbs:       float32(carbs),

		// 🔥 SHU YERDA PORTION YO‘QOLMAYDI
		Portion: portion,
	}

	// =========================
	// CALL AI
	// =========================

	res, err := u.aiClient.AnalyzeMeal(aiReq)
	if err != nil {
		return nil, err
	}

	return res, nil
}