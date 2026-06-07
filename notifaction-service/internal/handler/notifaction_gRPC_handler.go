package handler

import (
	"context"
	"notifaction-service/internal/models"
	"notifaction-service/internal/usecase"

	pb "github.com/khbdev/what-food-proto/proto/notifaction-crud"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	uc *usecase.NotificationUsecase
}

func NewNotificationHandler(uc *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		uc: uc,
	}
}

func (h *NotificationHandler) CreateNotification(
	ctx context.Context,
	req *pb.CreateNotificationRequest,
) (*pb.NotificationResponse, error) {

	n := &models.Notification{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
	}

	if err := h.uc.Create(ctx, n); err != nil {
		return nil, err
	}

	return &pb.NotificationResponse{
		Notification: mapToProto(n),
	}, nil
}

func (h *NotificationHandler) GetNotification(
	ctx context.Context,
	req *pb.GetNotificationRequest,
) (*pb.NotificationResponse, error) {

	n, err := h.uc.GetByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.NotificationResponse{
		Notification: mapToProto(n),
	}, nil
}

func (h *NotificationHandler) GetNotifications(
	ctx context.Context,
	req *pb.GetNotificationsRequest,
) (*pb.GetNotificationsResponse, error) {

	list, err := h.uc.GetAll(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Notification, 0, len(list))

	for _, n := range list {
		res = append(res, mapToProto(&n))
	}

	return &pb.GetNotificationsResponse{
		Notifications: res,
	}, nil
}

func (h *NotificationHandler) UpdateNotification(
	ctx context.Context,
	req *pb.UpdateNotificationRequest,
) (*pb.NotificationResponse, error) {

	n := &models.Notification{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageUrl,
	}

	if err := h.uc.Update(n); err != nil {
		return nil, err
	}

	return &pb.NotificationResponse{
		Notification: mapToProto(n),
	}, nil
}

func (h *NotificationHandler) DeleteNotification(
	ctx context.Context,
	req *pb.DeleteNotificationRequest,
) (*pb.DeleteNotificationResponse, error) {

	if err := h.uc.Delete(req.Id); err != nil {
		return nil, err
	}

	return &pb.DeleteNotificationResponse{
		Success: true,
	}, nil
}

func mapToProto(n *models.Notification) *pb.Notification {
	return &pb.Notification{
		Id:          n.ID,
		Title:       n.Title,
		Description: n.Description,
		ImageUrl:    n.ImageURL,
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
	}
}
