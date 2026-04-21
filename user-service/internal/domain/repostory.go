package domain

import (
	"context"

	"your_project/internal/models"
)

// UserRepository — User uchun CRUD contract (interface)
type UserRepository interface {
	// Create yangi user yaratadi
	Create(ctx context.Context, user *models.User) error

	// GetByID ID bo‘yicha user olish
	GetByID(ctx context.Context, id uint) (*models.User, error)

	// GetByPhone phone orqali user topish
	GetByPhone(ctx context.Context, phone string) (*models.User, error)

	// GetAll barcha userlarni olish (pagination bilan keyin kengaytirasan)
	GetAll(ctx context.Context) ([]models.User, error)

	// Update user ma’lumotlarini yangilash
	Update(ctx context.Context, user *models.User) error

	// Delete user o‘chirish (soft delete bo‘lishi mumkin)
	Delete(ctx context.Context, id uint) error
}