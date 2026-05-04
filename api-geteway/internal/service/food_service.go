package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

type FoodService struct {
	client *client.FoodClient
}

// =========================
// INIT
// =========================

func NewFoodService(c *client.FoodClient) *FoodService {
	return &FoodService{client: c}
}

// =========================
// RECIPE
// =========================

func (s *FoodService) CreateRecipe(ctx context.Context, req *foodpb.CreateRecipeRequest) error {

	if req.Recipe == nil {
		return errors.New("recipe is required")
	}

	if req.Recipe.RestaurantId == 0 {
		return errors.New("restaurant_id is required")
	}

	if strings.TrimSpace(req.Recipe.Name) == "" {
		return errors.New("name is required")
	}

	return s.client.CreateRecipe(req)
}

func (s *FoodService) GetRecipeByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Recipe, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.client.GetRecipeByID(req.Id)
}

func (s *FoodService) GetAllRecipes(ctx context.Context) ([]*foodpb.Recipe, error) {
	return s.client.GetAllRecipes()
}

func (s *FoodService) UpdateRecipe(ctx context.Context, req *foodpb.UpdateRecipeRequest) error {

	if req.Recipe == nil {
		return errors.New("recipe is required")
	}

	if req.Recipe.Id == 0 {
		return errors.New("id is required")
	}

	if strings.TrimSpace(req.Recipe.Name) == "" {
		return errors.New("name cannot be empty")
	}

	return s.client.UpdateRecipe(req)
}

func (s *FoodService) DeleteRecipe(ctx context.Context, req *foodpb.GetByIDRequest) error {

	if req.Id == 0 {
		return errors.New("id is required")
	}

	return s.client.DeleteRecipe(req.Id)
}

// =========================
// SALAD
// =========================

func (s *FoodService) CreateSalad(ctx context.Context, req *foodpb.CreateSaladRequest) error {

	if req.Salad == nil {
		return errors.New("salad is required")
	}

	if strings.TrimSpace(req.Salad.Name) == "" {
		return errors.New("name is required")
	}

	return s.client.CreateSalad(req)
}

func (s *FoodService) GetSaladByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Salad, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.client.GetSaladByID(req.Id)
}

func (s *FoodService) GetAllSalads(ctx context.Context) ([]*foodpb.Salad, error) {
	return s.client.GetAllSalads()
}

func (s *FoodService) UpdateSalad(ctx context.Context, req *foodpb.UpdateSaladRequest) error {

	if req.Salad == nil {
		return errors.New("salad is required")
	}

	if req.Salad.Id == 0 {
		return errors.New("id is required")
	}

	if strings.TrimSpace(req.Salad.Name) == "" {
		return errors.New("name cannot be empty")
	}

	return s.client.UpdateSalad(req)
}

func (s *FoodService) DeleteSalad(ctx context.Context, req *foodpb.GetByIDRequest) error {

	if req.Id == 0 {
		return errors.New("id is required")
	}

	return s.client.DeleteSalad(req.Id)
}

// =========================
// RESTAURANT
// =========================

func (s *FoodService) CreateRestaurant(ctx context.Context, req *foodpb.CreateRestaurantRequest) (int64, error) {

	if strings.TrimSpace(req.RestaurantName) == "" {
		return 0, errors.New("restaurant_name is required")
	}

	return s.client.CreateRestaurant(req)
}

func (s *FoodService) GetRestaurantByID(ctx context.Context, req *foodpb.GetByIDRequest) (*foodpb.Restaurant, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.client.GetRestaurantByID(req.Id)
}

func (s *FoodService) GetAllRestaurants(ctx context.Context) ([]*foodpb.Restaurant, error) {
	return s.client.GetAllRestaurants()
}

func (s *FoodService) UpdateRestaurant(ctx context.Context, req *foodpb.UpdateRestaurantRequest) error {

	if req.Restaurant == nil {
		return errors.New("restaurant is required")
	}

	if req.Restaurant.Id == 0 {
		return errors.New("id is required")
	}

	if strings.TrimSpace(req.Restaurant.RestaurantName) == "" {
		return errors.New("restaurant_name cannot be empty")
	}

	return s.client.UpdateRestaurant(req)
}

func (s *FoodService) DeleteRestaurant(ctx context.Context, req *foodpb.GetByIDRequest) error {

	if req.Id == 0 {
		return errors.New("id is required")
	}

	return s.client.DeleteRestaurant(req.Id)
}

// =========================
// FILTER
// =========================

func (s *FoodService) FilterFood(ctx context.Context, req *foodpb.FoodFilterRequest) ([]*foodpb.FoodItem, error) {

	// optional validation (soft rules)
	if req.MaxKcal < 0 {
		return nil, errors.New("max_kcal cannot be negative")
	}

	return s.client.FilterFood(req)
}