package repository

import (
	"notifaction-service/internal/domain"
	"notifaction-service/internal/models"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) Create(n *models.Notification) error {
	return r.db.Create(n).Error
}

func (r *notificationRepository) GetByID(id string) (*models.Notification, error) {
	var notification models.Notification

	err := r.db.First(&notification, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) GetAll(limit, offset int) ([]models.Notification, error) {
	var notifications []models.Notification

	err := r.db.
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&notifications).Error

	return notifications, err
}

func (r *notificationRepository) Update(n *models.Notification) error {
	return r.db.Save(n).Error
}

func (r *notificationRepository) Delete(id string) error {
	return r.db.Delete(&models.Notification{}, "id = ?", id).Error
}
