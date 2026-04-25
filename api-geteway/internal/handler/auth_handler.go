package handler

import (
	"net/http"

	"api-geteway/internal/models"
	"api-geteway/pkg/response"

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
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.Register(c.Request.Context(), &authpb.RegisterRequest{
		FullName: req.FullName,
		Phone:    req.Phone,
		Age:      req.Age,
		Address:  req.Address,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}
// =========================
// LOGIN
// =========================
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.Login(c.Request.Context(), &authpb.LoginRequest{
		Phone: req.Phone,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// VERIFY OTP
// =========================

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req models.VerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "invalid request")
		return
	}

	res, err := h.svc.VerifyOTP(c.Request.Context(), &authpb.VerifyRequest{
		Otp: req.OTP,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, "otp verification failed")
		return
	}

	response.OK(c, res)
}