package handler

import (
	"context"

	"food-service/internal/domain"
	"food-service/internal/models"

	foodpb "github.com/khbdev/what-food-proto/proto/auth"
)

type FoodHandler struct {
	foodpb.UnimplementedRecipeServiceServer
	foodpb.UnimplementedSaladServiceServer
	foodpb.UnimplementedRestaurantServiceServer

	recipeUC     domain.RecipeUsecase
	saladUC      domain.SaladUsecase
	restaurantUC domain.RestaurantUsecase
}

func NewFoodHandler(
	r domain.RecipeUsecase,
	s domain.SaladUsecase,
	res domain.RestaurantUsecase,
) *FoodHandler {
	return &FoodHandler{
		recipeUC:     r,
		saladUC:      s,
		restaurantUC: res,
	}
}

// ===== RECIPE =====

func (h *FoodHandler) CreateRecipe(ctx context.Context, req *foodpb.CreateRecipeRequest) (*foodpb.Empty, error) {
	r := req.Recipe

	err := h.recipeUC.Create(ctx, &models.Recipe{
		RestaurantID: r.RestaurantId,
		Name:         r.Name,
		Description:  r.Description,
		ImageURL:     r.ImageUrl,
		VideoURL:     r.VideoUrl,
		Country:      r.Country,
		MealTime:     r.MealTime,
		Kcal:         int(r.Kcal),
		Protein:      r.Protein,
		Fat:          r.Fat,
		Carbs:        r.Carbs,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.Empty{}, nil
}

func (h *FoodHandler) GetRecipeByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.RecipeResponse, error) {
	res, err := h.recipeUC.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &foodpb.RecipeResponse{
		Recipe: &foodpb.Recipe{
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

func (h *FoodHandler) GetAllRecipes(ctx context.Context, _ *foodpb.Empty) (*foodpb.RecipeListResponse, error) {
	list, err := h.recipeUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*foodpb.Recipe
	for _, r := range list {
		out = append(out, &foodpb.Recipe{
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

	return &foodpb.RecipeListResponse{Recipes: out}, nil
}

func (h *FoodHandler) UpdateRecipe(ctx context.Context, req *foodpb.UpdateRecipeRequest) (*foodpb.Empty, error) {
	r := req.Recipe

	err := h.recipeUC.Update(ctx, &models.Recipe{
		ID:           r.Id,
		RestaurantID: r.RestaurantId,
		Name:         r.Name,
		Description:  r.Description,
		ImageURL:     r.ImageUrl,
		VideoURL:     r.VideoUrl,
		Country:      r.Country,
		MealTime:     r.MealTime,
		Kcal:         int(r.Kcal),
		Protein:      r.Protein,
		Fat:          r.Fat,
		Carbs:        r.Carbs,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.Empty{}, nil
}

func (h *FoodHandler) DeleteRecipe(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Empty, error) {
	if err := h.recipeUC.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &foodpb.Empty{}, nil
}

// ===== SALAD =====

func (h *FoodHandler) CreateSalad(ctx context.Context, req *foodpb.CreateSaladRequest) (*foodpb.Empty, error) {
	s := req.Salad

	err := h.saladUC.Create(ctx, &models.Salad{
		RestaurantID: s.RestaurantId,
		Name:         s.Name,
		Description:  s.Description,
		ImageURL:     s.ImageUrl,
		VideoURL:     s.VideoUrl,
		Country:      s.Country,
		MealTime:     s.MealTime,
		Kcal:         int(s.Kcal),
		Protein:      s.Protein,
		Fat:          s.Fat,
		Carbs:        s.Carbs,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.Empty{}, nil
}

func (h *FoodHandler) GetSaladByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.SaladResponse, error) {
	res, err := h.saladUC.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &foodpb.SaladResponse{
		Salad: &foodpb.Salad{
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

func (h *FoodHandler) GetAllSalads(ctx context.Context, _ *foodpb.Empty) (*foodpb.SaladListResponse, error) {
	list, err := h.saladUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*foodpb.Salad
	for _, s := range list {
		out = append(out, &foodpb.Salad{
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

	return &foodpb.SaladListResponse{Salads: out}, nil
}

func (h *FoodHandler) UpdateSalad(ctx context.Context, req *foodpb.UpdateSaladRequest) (*foodpb.Empty, error) {
	s := req.Salad

	err := h.saladUC.Update(ctx, &models.Salad{
		ID:           s.Id,
		RestaurantID: s.RestaurantId,
		Name:         s.Name,
		Description:  s.Description,
		ImageURL:     s.ImageUrl,
		VideoURL:     s.VideoUrl,
		Country:      s.Country,
		MealTime:     s.MealTime,
		Kcal:         int(s.Kcal),
		Protein:      s.Protein,
		Fat:          s.Fat,
		Carbs:        s.Carbs,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.Empty{}, nil
}

func (h *FoodHandler) DeleteSalad(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Empty, error) {
	if err := h.saladUC.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &foodpb.Empty{}, nil
}

// ===== RESTAURANT =====

func (h *FoodHandler) CreateRestaurant(ctx context.Context, req *foodpb.CreateRestaurantRequest) (*foodpb.CreateRestaurantResponse, error) {
	id, err := h.restaurantUC.Create(ctx, &models.Restaurant{
		RestaurantName: req.RestaurantName,
		Description:    req.Description,
		ImageURL:       req.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.CreateRestaurantResponse{Id: id}, nil
}

func (h *FoodHandler) GetRestaurantByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.RestaurantResponse, error) {
	res, err := h.restaurantUC.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &foodpb.RestaurantResponse{
		Restaurant: &foodpb.Restaurant{
			Id:             res.ID,
			RestaurantName: res.RestaurantName,
			Description:    res.Description,
			ImageUrl:       res.ImageURL,
		},
	}, nil
}

func (h *FoodHandler) GetAllRestaurants(ctx context.Context, _ *foodpb.Empty) (*foodpb.RestaurantListResponse, error) {
	list, err := h.restaurantUC.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []*foodpb.Restaurant
	for _, r := range list {
		out = append(out, &foodpb.Restaurant{
			Id:             r.ID,
			RestaurantName: r.RestaurantName,
			Description:    r.Description,
			ImageUrl:       r.ImageURL,
		})
	}

	return &foodpb.RestaurantListResponse{Restaurants: out}, nil
}

func (h *FoodHandler) UpdateRestaurant(ctx context.Context, req *foodpb.UpdateRestaurantRequest) (*foodpb.Empty, error) {
	r := req.Restaurant

	err := h.restaurantUC.Update(ctx, &models.Restaurant{
		ID:             r.Id,
		RestaurantName: r.RestaurantName,
		Description:    r.Description,
		ImageURL:       r.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &foodpb.Empty{}, nil
}

func (h *FoodHandler) DeleteRestaurant(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Empty, error) {
	if err := h.restaurantUC.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &foodpb.Empty{}, nil
}