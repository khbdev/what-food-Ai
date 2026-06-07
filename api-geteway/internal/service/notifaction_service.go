package service

import (
	"api-geteway/internal/client"
	"context"
	"errors"
	"strings"

	notificationpb "github.com/khbdev/what-food-proto/proto/notifaction-crud"
)

type NotificationService struct {
	notificationClient *client.NotificationClient
}

func NewNotificationService(c *client.NotificationClient) *NotificationService {
	return &NotificationService{
		notificationClient: c,
	}
}

func (s *NotificationService) CreateNotification(
	ctx context.Context,
	req *notificationpb.CreateNotificationRequest,
) (*notificationpb.NotificationResponse, error) {

	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("title is required")
	}

	if strings.TrimSpace(req.Description) == "" {
		return nil, errors.New("description is required")
	}

	if len(req.Title) > 255 {
		return nil, errors.New("title is too long")
	}

	return s.notificationClient.CreateNotification(req)
}

func (s *NotificationService) GetNotification(
	ctx context.Context,
	req *notificationpb.GetNotificationRequest,
) (*notificationpb.NotificationResponse, error) {

	if strings.TrimSpace(req.Id) == "" {
		return nil, errors.New("id is required")
	}

	return s.notificationClient.GetNotification(req)
}

func (s *NotificationService) GetNotifications(
	ctx context.Context,
	req *notificationpb.GetNotificationsRequest,
) (*notificationpb.GetNotificationsResponse, error) {

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		return nil, errors.New("limit too large")
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	return s.notificationClient.GetNotifications(req)
}

func (s *NotificationService) UpdateNotification(
	ctx context.Context,
	req *notificationpb.UpdateNotificationRequest,
) (*notificationpb.NotificationResponse, error) {

	if strings.TrimSpace(req.Id) == "" {
		return nil, errors.New("id is required")
	}

	if req.Title != "" && strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("title cannot be empty")
	}

	if req.Description != "" && strings.TrimSpace(req.Description) == "" {
		return nil, errors.New("description cannot be empty")
	}

	return s.notificationClient.UpdateNotification(req)
}

func (s *NotificationService) DeleteNotification(
	ctx context.Context,
	req *notificationpb.DeleteNotificationRequest,
) (*notificationpb.DeleteNotificationResponse, error) {

	if strings.TrimSpace(req.Id) == "" {
		return nil, errors.New("id is required")
	}

	return s.notificationClient.DeleteNotification(req)
}
