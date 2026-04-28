package handler

import (
	"context"
	"time"

	"user-product-service/internal/usecase"
	"user-product-service/internal/models"

	incrideatspb "github.com/khbdev/what-food-proto/proto/incrideats"
)

type ProductsHandler struct {
	incrideatspb.UnimplementedIngredientServiceServer
	uc *usecase.IngredientUsecase
}

func NewProductsHandler(uc *usecase.IngredientUsecase) *ProductsHandler {
	return &ProductsHandler{
		uc: uc,
	}
}

// ===== CREATE =====

func (h *ProductsHandler) Create(ctx context.Context, req *incrideatspb.CreateIngredientRequest) (*incrideatspb.IngredientResponse, error) {
	input := usecase.IngredientInput{
		UserID:     req.GetUserId(),
		Name:       req.GetName(),
		Quantity:   req.GetQuantity(),
		CategoryID: req.GetCategoryId(),
	}

	res, err := h.uc.Create(input)
	if err != nil {
		return nil, err
	}

	return &incrideatspb.IngredientResponse{
		Ingredient: mapIngredient(res),
	}, nil
}

// ===== GET BY ID =====

func (h *ProductsHandler) GetByID(ctx context.Context, req *incrideatspb.GetIngredientByIDRequest) (*incrideatspb.IngredientResponse, error) {
	res, err := h.uc.GetByID(req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &incrideatspb.IngredientResponse{
		Ingredient: mapIngredient(res),
	}, nil
}

// ===== GET ALL =====

func (h *ProductsHandler) GetAll(ctx context.Context, req *incrideatspb.GetAllIngredientsRequest) (*incrideatspb.IngredientListResponse, error) {
	res, err := h.uc.GetAll(req.GetUserId())
	if err != nil {
		return nil, err
	}

	var items []*incrideatspb.Ingredient
	for _, i := range res {
		ing := i
		items = append(items, mapIngredient(&ing))
	}

	return &incrideatspb.IngredientListResponse{
		Ingredients: items,
	}, nil
}

// ===== UPDATE =====

func (h *ProductsHandler) Update(ctx context.Context, req *incrideatspb.UpdateIngredientRequest) (*incrideatspb.IngredientResponse, error) {
	input := usecase.IngredientInput{
		ID:         req.GetId(),
		UserID:     req.GetUserId(),
		Name:       req.GetName(),
		Quantity:   req.GetQuantity(),
		CategoryID: req.GetCategoryId(),
	}

	res, err := h.uc.Update(input)
	if err != nil {
		return nil, err
	}

	return &incrideatspb.IngredientResponse{
		Ingredient: mapIngredient(res),
	}, nil
}

// ===== DELETE =====

func (h *ProductsHandler) Delete(ctx context.Context, req *incrideatspb.DeleteIngredientRequest) (*incrideatspb.Empty, error) {
	err := h.uc.Delete(req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &incrideatspb.Empty{}, nil
}

// ===== MAPPER =====

func mapIngredient(m *models.Ingredient) *incrideatspb.Ingredient {
	return &incrideatspb.Ingredient{
		Id:         m.ID,
		UserId:     m.UserID,
		Name:       m.Name,
		Quantity:   m.Quantity,
		CategoryId: m.CategoryID,
		CreatedAt:  m.CreatedAt.Format(time.RFC3339),
	}
}