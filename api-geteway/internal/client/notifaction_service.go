package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	notificationpb "github.com/khbdev/what-food-proto/proto/notifaction-crud"
)

type NotificationClient struct {
	conn *grpc.ClientConn
	svc  notificationpb.NotificationServiceClient
}

const (
	timeoutNotification = 5 * time.Second
)

func NewNotificationClient(addr string) (*NotificationClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutNotification)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Notification service connected:", addr)

	return &NotificationClient{
		conn: conn,
		svc:  notificationpb.NewNotificationServiceClient(conn),
	}, nil
}

func (c *NotificationClient) Close() error {
	return c.conn.Close()
}

func (c *NotificationClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutNotification)
}

func (c *NotificationClient) CreateNotification(req *notificationpb.CreateNotificationRequest) (*notificationpb.NotificationResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.CreateNotification(ctx, req)
}

func (c *NotificationClient) GetNotification(req *notificationpb.GetNotificationRequest) (*notificationpb.NotificationResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetNotification(ctx, req)
}

func (c *NotificationClient) GetNotifications(req *notificationpb.GetNotificationsRequest) (*notificationpb.GetNotificationsResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetNotifications(ctx, req)
}

func (c *NotificationClient) UpdateNotification(req *notificationpb.UpdateNotificationRequest) (*notificationpb.NotificationResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.UpdateNotification(ctx, req)
}

func (c *NotificationClient) DeleteNotification(req *notificationpb.DeleteNotificationRequest) (*notificationpb.DeleteNotificationResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.DeleteNotification(ctx, req)
}
