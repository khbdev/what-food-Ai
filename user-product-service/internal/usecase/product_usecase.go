package usecase

import (
	"errors"
	"time"

	"user-product-service/internal/domain"
	"user-product-service/internal/models"
)

type IngredientUsecase struct {
	repo domain.IngredientRepository
}

func NewIngredientUsecase(repo domain.IngredientRepository) *IngredientUsecase {
	return &IngredientUsecase{
		repo: repo,
	}
}

type IngredientInput struct {
	ID         int64
	UserID     int64
	Name       string
	Quantity   float64
	CategoryID int64
}

func (u *IngredientUsecase) Create(input IngredientInput) (*models.Ingredient, error) {
	if input.UserID <= 0 {
		return nil, errors.New("invalid user_id")
	}

	if input.Name == "" {
		return nil, errors.New("name is required")
	}

	if input.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than 0")
	}

	if input.CategoryID <= 0 {
		return nil, errors.New("category_id is required")
	}

	ing := &models.Ingredient{
		UserID:     input.UserID,
		Name:       input.Name,
		Quantity:   input.Quantity,
		CategoryID: input.CategoryID,
		CreatedAt:  time.Now(),
	}

	if err := u.repo.Create(ing); err != nil {
		return nil, err
	}

	return ing, nil
}

func (u *IngredientUsecase) GetByID(id int64, userID int64) (*models.Ingredient, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	if userID <= 0 {
		return nil, errors.New("invalid user_id")
	}

	return u.repo.GetByID(id, userID)
}

func (u *IngredientUsecase) GetAll(userID int64) ([]models.Ingredient, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user_id")
	}

	return u.repo.GetAll(userID)
}

func (u *IngredientUsecase) Update(input IngredientInput) (*models.Ingredient, error) {
	if input.ID <= 0 {
		return nil, errors.New("invalid id")
	}

	if input.UserID <= 0 {
		return nil, errors.New("invalid user_id")
	}

	ing, err := u.repo.GetByID(input.ID, input.UserID)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		ing.Name = input.Name
	}

	if input.Quantity > 0 {
		ing.Quantity = input.Quantity
	}

	if input.CategoryID > 0 {
		ing.CategoryID = input.CategoryID
	}

	if err := u.repo.Update(ing); err != nil {
		return nil, err
	}

	return ing, nil
}

func (u *IngredientUsecase) Delete(id int64, userID int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	if userID <= 0 {
		return errors.New("invalid user_id")
	}

	return u.repo.Delete(id, userID)
}