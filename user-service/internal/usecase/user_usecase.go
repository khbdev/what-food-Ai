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

// Create — Write-through
func (u *UserUsecase) Create(ctx context.Context, req *models.User) error {
	user := &models.User{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("usecase.Create: %w", err)
	}

	_ = u.cache.SetUser(ctx, user, userCacheTTL)
	return nil
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

// GetAll — Read-through + pagination
func (u *UserUsecase) GetAll(ctx context.Context, req *models.GetAllUsersRequest) (*models.GetAllUsersResponse, error) {
	users, err := u.repo.GetAll(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetAll: %w", err)
	}

	for i := range users {
		_ = u.cache.SetUser(ctx, &users[i], userCacheTTL)
	}

	return &models.GetAllUsersResponse{
		Users:  users,
		Total:  len(users),
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}

// Update — Write-through
func (u *UserUsecase) Update(ctx context.Context, id uint, req *models.UpdateUserRequest) error {
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