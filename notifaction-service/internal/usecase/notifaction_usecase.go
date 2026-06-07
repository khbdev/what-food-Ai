package usecase

import (
	"context"
	"notifaction-service/internal/domain"
	"notifaction-service/internal/models"
)

type NotificationUsecase struct {
	repo      domain.NotificationRepository
	publisher domain.NotificationPublisher
}

func NewNotificationUsecase(
	repo domain.NotificationRepository,
	publisher domain.NotificationPublisher,
) *NotificationUsecase {
	return &NotificationUsecase{
		repo:      repo,
		publisher: publisher,
	}
}

func (u *NotificationUsecase) Create(ctx context.Context, n *models.Notification) error {
	// 1. DB ga saqlash
	if err := u.repo.Create(n); err != nil {
		return err
	}

	// 2. Redis event publish (REAL-TIME)
	if err := u.publisher.PublishCreated(ctx, n); err != nil {
		return err
	}

	return nil
}

func (u *NotificationUsecase) GetByID(id string) (*models.Notification, error) {
	return u.repo.GetByID(id)
}

func (u *NotificationUsecase) GetAll(limit, offset int) ([]models.Notification, error) {
	return u.repo.GetAll(limit, offset)
}

func (u *NotificationUsecase) Update(n *models.Notification) error {
	return u.repo.Update(n)
}

func (u *NotificationUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
