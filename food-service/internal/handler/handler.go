package handler

import (
	"context"

	"food-service/genproto/food"
	"food-service/internal/domain"
	"food-service/internal/models"
)

type GRPCHandler struct {
	food.UnimplementedRecipeServiceServer
	food.UnimplementedSaladServiceServer
	food.UnimplementedRestaurantServiceServer

	recipeUC     domain.RecipeUsecase
	saladUC      domain.SaladUsecase
	restaurantUC domain.RestaurantUsecase
}

func NewGRPCHandler(
	r domain.RecipeUsecase,
	s domain.SaladUsecase,
	res domain.RestaurantUsecase,
) *GRPCHandler {
	return &GRPCHandler{
		recipeUC:     r,
		saladUC:      s,
		restaurantUC: res,
	}
}

// ================= RECIPE =================

func (h *GRPCHandler) CreateRecipe(ctx context.Context, req *food.CreateRecipeRequest) (*food.Empty, error) {
	r := req.GetRecipe()

	model := &models.Recipe{
		RestaurantID: r.GetRestaurantId(),
		Name:         r.GetName(),
		Description:  r.GetDescription(),
		ImageURL:     r.GetImageUrl(),
		VideoURL:     r.GetVideoUrl(),
		Country:      r.GetCountry(),
		MealTime:     r.GetMealTime(),
		Kcal:         int(r.GetKcal()),
		Protein:      r.GetProtein(),
		Fat:          r.GetFat(),
		Carbs:        r.GetCarbs(),
	}

	if err := h.recipeUC.Create(ctx, model); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

func (h *GRPCHandler) GetRecipeByID(ctx context.Context, req *food.GetByIDRequest) (*food.RecipeResponse, error) {
	res, err := h.recipeUC.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &food.RecipeResponse{
		Recipe: &food.Recipe{
			Id:           res.ID,
			RestaurantId: res.RestaurantID,
			Name:         res.Name,
			Description:  res.Description,
			ImageUrl:     res.ImageURL,
			VideoUrl:     res.VideoURL,
			Country:      res.Country,
			MealTime:     res.MealTime,
			Kcal:         int32(res.Kcal),
			Protein:      res.Protein,
			Fat:          res.Fat,
			Carbs:        res.Carbs,
		},
	}, nil
}

func (h *GRPCHandler) GetAllRecipes(ctx context.Context, _ *food.Empty) (*food.RecipeListResponse, error) {
	list, err := h.recipeUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*food.Recipe
	for _, r := range list {
		out = append(out, &food.Recipe{
			Id:           r.ID,
			RestaurantId: r.RestaurantID,
			Name:         r.Name,
			Description:  r.Description,
			ImageUrl:     r.ImageURL,
			VideoUrl:     r.VideoURL,
			Country:      r.Country,
			MealTime:     r.MealTime,
			Kcal:         int32(r.Kcal),
			Protein:      r.Protein,
			Fat:          r.Fat,
			Carbs:        r.Carbs,
		})
	}

	return &food.RecipeListResponse{Recipes: out}, nil
}

func (h *GRPCHandler) UpdateRecipe(ctx context.Context, req *food.UpdateRecipeRequest) (*food.Empty, error) {
	r := req.GetRecipe()

	model := &models.Recipe{
		ID:           r.GetId(),
		RestaurantID: r.GetRestaurantId(),
		Name:         r.GetName(),
		Description:  r.GetDescription(),
		ImageURL:     r.GetImageUrl(),
		VideoURL:     r.GetVideoUrl(),
		Country:      r.GetCountry(),
		MealTime:     r.GetMealTime(),
		Kcal:         int(r.GetKcal()),
		Protein:      r.GetProtein(),
		Fat:          r.GetFat(),
		Carbs:        r.GetCarbs(),
	}

	if err := h.recipeUC.Update(ctx, model); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

func (h *GRPCHandler) DeleteRecipe(ctx context.Context, req *food.GetByIDRequest) (*food.Empty, error) {
	if err := h.recipeUC.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

// ================= SALAD =================

func (h *GRPCHandler) CreateSalad(ctx context.Context, req *food.CreateSaladRequest) (*food.Empty, error) {
	s := req.GetSalad()

	model := &models.Salad{
		RestaurantID: s.GetRestaurantId(),
		Name:         s.GetName(),
		Description:  s.GetDescription(),
		ImageURL:     s.GetImageUrl(),
		VideoURL:     s.GetVideoUrl(),
		Country:      s.GetCountry(),
		MealTime:     s.GetMealTime(),
		Kcal:         int(s.GetKcal()),
		Protein:      s.GetProtein(),
		Fat:          s.GetFat(),
		Carbs:        s.GetCarbs(),
	}

	if err := h.saladUC.Create(ctx, model); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

func (h *GRPCHandler) GetSaladByID(ctx context.Context, req *food.GetByIDRequest) (*food.SaladResponse, error) {
	res, err := h.saladUC.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &food.SaladResponse{
		Salad: &food.Salad{
			Id:           res.ID,
			RestaurantId: res.RestaurantID,
			Name:         res.Name,
			Description:  res.Description,
			ImageUrl:     res.ImageURL,
			VideoUrl:     res.VideoURL,
			Country:      res.Country,
			MealTime:     res.MealTime,
			Kcal:         int32(res.Kcal),
			Protein:      res.Protein,
			Fat:          res.Fat,
			Carbs:        res.Carbs,
		},
	}, nil
}

func (h *GRPCHandler) GetAllSalads(ctx context.Context, _ *food.Empty) (*food.SaladListResponse, error) {
	list, err := h.saladUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*food.Salad
	for _, s := range list {
		out = append(out, &food.Salad{
			Id:           s.ID,
			RestaurantId: s.RestaurantID,
			Name:         s.Name,
			Description:  s.Description,
			ImageUrl:     s.ImageURL,
			VideoUrl:     s.VideoURL,
			Country:      s.Country,
			MealTime:     s.MealTime,
			Kcal:         int32(s.Kcal),
			Protein:      s.Protein,
			Fat:          s.Fat,
			Carbs:        s.Carbs,
		})
	}

	return &food.SaladListResponse{Salads: out}, nil
}

func (h *GRPCHandler) UpdateSalad(ctx context.Context, req *food.UpdateSaladRequest) (*food.Empty, error) {
	s := req.GetSalad()

	model := &models.Salad{
		ID:           s.GetId(),
		RestaurantID: s.GetRestaurantId(),
		Name:         s.GetName(),
		Description:  s.GetDescription(),
		ImageURL:     s.GetImageUrl(),
		VideoURL:     s.GetVideoUrl(),
		Country:      s.GetCountry(),
		MealTime:     s.GetMealTime(),
		Kcal:         int(s.GetKcal()),
		Protein:      s.GetProtein(),
		Fat:          s.GetFat(),
		Carbs:        s.GetCarbs(),
	}

	if err := h.saladUC.Update(ctx, model); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

func (h *GRPCHandler) DeleteSalad(ctx context.Context, req *food.GetByIDRequest) (*food.Empty, error) {
	if err := h.saladUC.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

// ================= RESTAURANT =================

func (h *GRPCHandler) CreateRestaurant(ctx context.Context, req *food.CreateRestaurantRequest) (*food.CreateRestaurantResponse, error) {
	id, err := h.restaurantUC.Create(ctx, &models.Restaurant{
		RestaurantName: req.GetRestaurantName(),
		Description:    req.GetDescription(),
		ImageURL:       req.GetImageUrl(),
	})
	if err != nil {
		return nil, err
	}

	return &food.CreateRestaurantResponse{Id: id}, nil
}

func (h *GRPCHandler) GetRestaurantByID(ctx context.Context, req *food.GetByIDRequest) (*food.RestaurantResponse, error) {
	res, err := h.restaurantUC.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &food.RestaurantResponse{
		Restaurant: &food.Restaurant{
			Id:             res.ID,
			RestaurantName: res.RestaurantName,
			Description:    res.Description,
			ImageUrl:       res.ImageURL,
		},
	}, nil
}

func (h *GRPCHandler) GetAllRestaurants(ctx context.Context, _ *food.Empty) (*food.RestaurantListResponse, error) {
	list, err := h.restaurantUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*food.Restaurant
	for _, r := range list {
		out = append(out, &food.Restaurant{
			Id:             r.ID,
			RestaurantName: r.RestaurantName,
			Description:    r.Description,
			ImageUrl:       r.ImageURL,
		})
	}

	return &food.RestaurantListResponse{Restaurants: out}, nil
}

func (h *GRPCHandler) UpdateRestaurant(ctx context.Context, req *food.UpdateRestaurantRequest) (*food.Empty, error) {
	r := req.GetRestaurant()

	model := &models.Restaurant{
		ID:             r.GetId(),
		RestaurantName: r.GetRestaurantName(),
		Description:    r.GetDescription(),
		ImageURL:       r.GetImageUrl(),
	}

	if err := h.restaurantUC.Update(ctx, model); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}

func (h *GRPCHandler) DeleteRestaurant(ctx context.Context, req *food.GetByIDRequest) (*food.Empty, error) {
	if err := h.restaurantUC.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &food.Empty{}, nil
}