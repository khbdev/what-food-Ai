package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	categorypb "github.com/khbdev/what-food-proto/proto/categorypb"
)

type CategoryService struct {
	client *client.UserProductClient
}

// =========================
// INIT
// =========================

func NewCategoryService(c *client.UserProductClient) *CategoryService {
	return &CategoryService{client: c}
}

// =========================
// CREATE CATEGORY
// =========================

func (s *CategoryService) CreateCategory(ctx context.Context, req *categorypb.CreateCategoryRequest) (*categorypb.CategoryResponse, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if len(req.Name) < 2 {
		return nil, errors.New("name too short")
	}

	return s.client.CreateCategory(req)
}

// =========================
// GET BY ID
// =========================

func (s *CategoryService) GetCategoryByID(ctx context.Context, req *categorypb.GetByIDRequest) (*categorypb.CategoryWithIngredientsResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.client.GetCategoryByID(req)
}

// =========================
// GET ALL
// =========================

func (s *CategoryService) GetAllCategories(ctx context.Context, req *categorypb.GetAllRequest) (*categorypb.CategoryListResponse, error) {

	return s.client.GetAllCategories(req)
}

// =========================
// UPDATE CATEGORY
// =========================

func (s *CategoryService) UpdateCategory(ctx context.Context, req *categorypb.UpdateCategoryRequest) (*categorypb.CategoryResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if req.Name != "" && len(req.Name) < 2 {
		return nil, errors.New("name too short")
	}

	return s.client.UpdateCategory(req)
}

// =========================
// DELETE CATEGORY
// =========================

func (s *CategoryService) DeleteCategory(ctx context.Context, req *categorypb.DeleteCategoryRequest) (*categorypb.Empty, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.client.DeleteCategory(req)
}

// =========================
// GET ALL WITH USER PRODUCTS
// =========================

func (s *CategoryService) GetAllWithUserProducts(ctx context.Context, req *categorypb.GetAllWithUserProductsRequest) (*categorypb.CategoryWithIngredientsListResponse, error) {

	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.client.GetAllWithUserProducts(req)
}