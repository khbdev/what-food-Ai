package domain

import "notifaction-service/internal/models"

type NotificationRepository interface {
	Create(n *models.Notification) error
	GetByID(id string) (*models.Notification, error)
	GetAll(limit, offset int) ([]models.Notification, error)
	Update(n *models.Notification) error
	Delete(id string) error
}
