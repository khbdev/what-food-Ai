package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	ingredientpb "github.com/khbdev/what-food-proto/proto/incrideatspb"
)

type ProductService struct {
	client *client.UserProductClient
}

// =========================
// INIT
// =========================

func NewProductService(c *client.UserProductClient) *ProductService {
	return &ProductService{client: c}
}

// =========================
// CREATE INGREDIENT
// =========================

func (s *ProductService) CreateIngredient(ctx context.Context, req *ingredientpb.CreateIngredientRequest) (*ingredientpb.IngredientResponse, error) {

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if req.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	if req.CategoryId == 0 {
		return nil, errors.New("category_id is required")
	}

	return s.client.CreateIngredient(req)
}

// =========================
// GET BY ID
// =========================

func (s *ProductService) GetIngredientByID(ctx context.Context, req *ingredientpb.GetIngredientByIDRequest) (*ingredientpb.IngredientResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.client.GetIngredientByID(req)
}

// =========================
// GET ALL
// =========================

func (s *ProductService) GetAllIngredients(ctx context.Context, req *ingredientpb.GetAllIngredientsRequest) (*ingredientpb.IngredientListResponse, error) {

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.client.GetAllIngredients(req)
}

// =========================
// UPDATE INGREDIENT
// =========================

func (s *ProductService) UpdateIngredient(ctx context.Context, req *ingredientpb.UpdateIngredientRequest) (*ingredientpb.IngredientResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if req.Quantity != 0 && req.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	if req.CategoryId != 0 && req.CategoryId <= 0 {
		return nil, errors.New("invalid category_id")
	}

	return s.client.UpdateIngredient(req)
}

// =========================
// DELETE INGREDIENT
// =========================

func (s *ProductService) DeleteIngredient(ctx context.Context, req *ingredientpb.DeleteIngredientRequest) (*ingredientpb.Empty, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.client.DeleteIngredient(req)
}