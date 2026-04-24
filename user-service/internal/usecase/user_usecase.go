package usecase

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/domain"
	"user-service/internal/models"
)

const userCacheTTL = 5 * time.Minute

type UserUsecase struct {
	repo  domain.UserRepository
	cache domain.UserCache
}

func NewUserUsecase(repo domain.UserRepository, cache domain.UserCache) *UserUsecase {
	return &UserUsecase{repo: repo, cache: cache}
}

// Create — Write-through, yaratilgan user'ni qaytaradi
func (u *UserUsecase) Create(ctx context.Context, req *models.User) (*models.User, error) {
	user := &models.User{
		Name:    req.Name,
		Phone:   req.Phone,
		Age:     req.Age,
		Address: req.Address,
		Email:   req.Email,
		Image:   req.Image,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("usecase.Create: %w", err)
	}

	// GORM Create dan keyin user.ID avtomatik to'ldiriladi
	_ = u.cache.SetUser(ctx, user, userCacheTTL)
	return user, nil
}

// GetByID — Read-through
func (u *UserUsecase) GetByID(ctx context.Context, id uint) (*models.User, error) {
	cached, err := u.cache.GetUser(ctx, id)
	if err == nil && cached != nil {
		return cached, nil
	}

	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetByID: %w", err)
	}

	_ = u.cache.SetUser(ctx, user, userCacheTTL)
	return user, nil
}

// GetByPhone — Read-through
func (u *UserUsecase) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	user, err := u.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetByPhone: %w", err)
	}

	_ = u.cache.SetUser(ctx, user, userCacheTTL)
	return user, nil
}

// GetAll — pagination
func (u *UserUsecase) GetAll(ctx context.Context, req *models.GetAllUsersRequest) (*models.GetAllUsersResponse, error) {
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	users, err := u.repo.GetAll(ctx, limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetAll: %w", err)
	}

	return &models.GetAllUsersResponse{
		Users:  users,
		Total:  len(users),
		Limit:  limit,
		Offset: req.Offset,
	}, nil
}

// Update — Write-through
func (u *UserUsecase) Update(ctx context.Context, id uint, req *models.User) error {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("usecase.Update: %w", err)
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.Image != "" {
		user.Image = req.Image
	}

	if err := u.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("usecase.Update: %w", err)
	}

	_ = u.cache.SetUser(ctx, user, userCacheTTL)
	return nil
}

// Delete — Cache invalidation
func (u *UserUsecase) Delete(ctx context.Context, id uint) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("usecase.Delete: %w", err)
	}

	_ = u.cache.DeleteUser(ctx, id)
	return nil
}