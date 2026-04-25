package handler

import (
	"errors"
	"net/http"

	"api-geteway/internal/models"
	"api-geteway/internal/service"
	"api-geteway/pkg/response"

	userrpb "github.com/khbdev/what-food-proto/proto/userr"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

// =========================
// INIT
// =========================

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

// =========================
// CREATE USER
// =========================

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.CreateUser(c.Request.Context(), &userrpb.CreateUserRequest{
		Name:    req.Name,
		Phone:   req.Phone,
		Age:     req.Age,
		Address: req.Address,
		Email:   req.Email,
		Image:   req.Image,
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

func (h *UserHandler) GetUserByID(c *gin.Context) {

	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id, err := str.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := h.svc.GetUserByID(c.Request.Context(), &userrpb.GetUserByIDRequest{
		Id: id,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET BY PHONE
// =========================

func (h *UserHandler) GetUserByPhone(c *gin.Context) {
	var req models.GetUserByPhoneRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.GetUserByPhone(c.Request.Context(), &userrpb.GetUserByPhoneRequest{
		Phone: req.Phone,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// GET ALL USERS
// =========================

func (h *UserHandler) GetAllUsers(c *gin.Context) {

	res, err := h.svc.GetAllUsers(c.Request.Context(), &userrpb.GetAllUsersRequest{})
	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// UPDATE USER
// =========================

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req models.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.UpdateUser(c.Request.Context(), &userrpb.UpdateUserRequest{
		Id:      req.ID,
		Name:    req.Name,
		Phone:   req.Phone,
		Age:     req.Age,
		Address: req.Address,
		Email:   req.Email,
		Image:   req.Image,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}

// =========================
// DELETE USER
// =========================

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var req models.DeleteUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.svc.DeleteUser(c.Request.Context(), &userrpb.DeleteUserRequest{
		Id: req.ID,
	})

	if err != nil {
		response.Fail(c, http.StatusBadRequest, err)
		return
	}

	response.OK(c, res)
}