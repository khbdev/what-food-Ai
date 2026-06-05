package handler

import (
	"net/http"

	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	svc *service.DashboardService
}

// =========================
// INIT
// =========================

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{svc: s}
}


func (h *DashboardHandler) GetDashboardStats(c *gin.Context) {

	res, err := h.svc.GetDashboardStats(c.Request.Context())
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}