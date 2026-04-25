package handler

import (
	"net/http"

	"api-geteway/internal/models"
	"api-geteway/internal/service"

	authpb "github.com/khbdev/what-food-proto/proto/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

// =========================
// REGISTER
// =========================

func (h *AuthHandler) Register(c *gin.Context) {

	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 🔥 MAP: models → proto
	protoReq := &authpb.RegisterRequest{
		FullName: req.FullName,
		Phone:    req.Phone,
		Age:      req.Age,
		Address:  req.Address,
	}

	res, err := h.svc.Register(c.Request.Context(), protoReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// =========================
// LOGIN
// =========================

func (h *AuthHandler) Login(c *gin.Context) {

	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 🔥 MAP
	protoReq := &authpb.LoginRequest{
		Phone: req.Phone,
	}

	res, err := h.svc.Login(c.Request.Context(), protoReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// =========================
// VERIFY OTP
// =========================

func (h *AuthHandler) VerifyOTP(c *gin.Context) {

	var req models.VerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 🔥 MAP
	protoReq := &authpb.VerifyRequest{
		Otp: req.OTP,
	}

	res, err := h.svc.VerifyOTP(c.Request.Context(), protoReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}