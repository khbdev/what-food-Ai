func (h *NutritionHandler) GetWeeklyNutrition(c *gin.Context) {

	// =========================
	// BODY (faqat optional filter)
	// =========================
	var req models.WeeklyNutritionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// USER ID FROM JWT CONTEXT
	// =========================
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	uid, ok := userID.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	// =========================
	// CALL SERVICE
	// =========================
	res, err := h.service.GetWeeklyNutrition(
		c.Request.Context(),
		&nutritionpb.WeeklyNutritionRequest{
			UserId: uid, // 🔥 FIX HERE
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// =========================
	// MAPPING RESPONSE
	// =========================
	days := make([]models.DailyNutrition, 0, len(res.Days))

	for _, d := range res.Days {
		days = append(days, models.DailyNutrition{
			Day:     d.Day,
			Kcal:    d.Kcal,
			Fat:     d.Fat,
			Carbs:   d.Carbs,
			Protein: d.Protein,
		})
	}

	// =========================
	// RESPONSE
	// =========================
	c.JSON(http.StatusOK, models.WeeklyNutritionResponse{
		Days: days,
	})
}