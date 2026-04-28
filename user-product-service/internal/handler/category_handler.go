package handler

import (
	"context"
	"time"

	"user-product-service/internal/usecase"
	"user-product-service/internal/models"

	categorypb "github.com/khbdev/what-food-proto/proto/products"
)

type CategoryHandler struct {
	categorypb.UnimplementedCategoryServiceServer
	uc *usecase.CategoryUsecase
}

func NewCategoryHandler(uc *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		uc: uc,
	}
}

// ===== CREATE =====

func (h *CategoryHandler) Create(ctx context.Context, req *categorypb.CreateCategoryRequest) (*categorypb.CategoryResponse, error) {
	res, err := h.uc.Create(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &categorypb.CategoryResponse{
		Category: mapCategory(res),
	}, nil
}

// ===== GET BY ID =====

func (h *CategoryHandler) GetByID(ctx context.Context, req *categorypb.GetByIDRequest) (*categorypb.CategoryWithIngredientsResponse, error) {
	res, err := h.uc.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &categorypb.CategoryWithIngredientsResponse{
		Category: mapCategoryWithIngredients(res),
	}, nil
}

// ===== GET ALL =====

func (h *CategoryHandler) GetAll(ctx context.Context, req *categorypb.GetAllRequest) (*categorypb.CategoryListResponse, error) {
	res, err := h.uc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var categories []*categorypb.Category
	for _, c := range res {
		cat := c
		categories = append(categories, mapCategory(&cat))
	}

	return &categorypb.CategoryListResponse{
		Categories: categories,
	}, nil
}

// ===== GET ALL WITH USER PRODUCTS =====

func (h *CategoryHandler) GetAllWithUserProducts(ctx context.Context, req *categorypb.GetAllWithUserProductsRequest) (*categorypb.CategoryWithIngredientsListResponse, error) {
	res, err := h.uc.GetAllWithUserProducts(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	var categories []*categorypb.CategoryWithIngredients
	for _, c := range res {
		cat := c
		categories = append(categories, mapCategoryWithIngredients(&cat))
	}

	return &categorypb.CategoryWithIngredientsListResponse{
		Categories: categories,
	}, nil
}

// ===== UPDATE =====

func (h *CategoryHandler) Update(ctx context.Context, req *categorypb.UpdateCategoryRequest) (*categorypb.CategoryResponse, error) {
	err := h.uc.Update(ctx, req.GetId(), req.GetName())
	if err != nil {
		return nil, err
	}

	// qayta olish (fresh data)
	res, err := h.uc.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &categorypb.CategoryResponse{
		Category: &categorypb.Category{
			Id:        res.CategoryID,
			Name:      res.Name,
			CreatedAt: time.Now().Format(time.RFC3339), // approximate
		},
	}, nil
}

// ===== DELETE =====

func (h *CategoryHandler) Delete(ctx context.Context, req *categorypb.DeleteCategoryRequest) (*categorypb.Empty, error) {
	if err := h.uc.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &categorypb.Empty{}, nil
}

// ===== MAPPERS =====

func mapCategory(m *models.Category) *categorypb.Category {
	return &categorypb.Category{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
	}
}

func mapCategoryWithIngredients(m *models.CategoryWithIngredients) *categorypb.CategoryWithIngredients {
	var items []*categorypb.Ingredient

	for _, i := range m.Items {
		ing := i
		items = append(items, &categorypb.Ingredient{
			Id:       ing.ID,
			Name:     ing.Name,
			Quantity: ing.Quantity,
		})
	}

	return &categorypb.CategoryWithIngredients{
		CategoryId: m.CategoryID,
		Name:       m.Name,
		Items:      items,
	}
}