package domain

import (
	"context"
	"notifaction-service/internal/models"
)

type NotificationPublisher interface {
	PublishCreated(ctx context.Context, n *models.Notification) error
}
