package redis_pub

import (
	"context"
	"encoding/json"
	"notifaction-service/internal/domain"
	"notifaction-service/internal/models"

	"github.com/redis/go-redis/v9"
)

type notificationPublisher struct {
	rdb *redis.Client
}

func NewNotificationPublisher(rdb *redis.Client) domain.NotificationPublisher {
	return &notificationPublisher{
		rdb: rdb,
	}
}

func (p *notificationPublisher) PublishCreated(ctx context.Context, n *models.Notification) error {
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return p.rdb.Publish(
		ctx,
		"notification.created",
		data,
	).Err()
}
