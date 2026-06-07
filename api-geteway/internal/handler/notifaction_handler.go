package handler

import (
	"api-geteway/internal/models"
	"api-geteway/internal/service"
	"api-geteway/pkg/response"
	"errors"
	"net/http"
	"strconv"

	notificationpb "github.com/khbdev/what-food-proto/proto/notifaction-crud"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	svc *service.NotificationService
}

// =========================
// INIT
// =========================

func NewNotificationHandler(s *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{svc: s}
}

// =========================
// CREATE NOTIFICATION
// =========================

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req models.CreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.CreateNotification(c.Request.Context(), &notificationpb.CreateNotificationRequest{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET BY ID
// =========================

func (h *NotificationHandler) GetNotification(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	res, err := h.svc.GetNotification(c.Request.Context(), &notificationpb.GetNotificationRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET ALL
// =========================

func (h *NotificationHandler) GetNotifications(c *gin.Context) {

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid limit"))
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid offset"))
		return
	}

	res, err := h.svc.GetNotifications(c.Request.Context(), &notificationpb.GetNotificationsRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// UPDATE
// =========================

func (h *NotificationHandler) UpdateNotification(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	var req models.UpdateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.UpdateNotification(c.Request.Context(), &notificationpb.UpdateNotificationRequest{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// DELETE
// =========================

func (h *NotificationHandler) DeleteNotification(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	res, err := h.svc.DeleteNotification(c.Request.Context(), &notificationpb.DeleteNotificationRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}
