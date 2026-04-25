package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	userrpb "github.com/khbdev/what-food-proto/proto/"
)

type UserService struct {
	userClient *client.UserClient
}

// =========================
// INIT
// =========================

func NewUserService(c *client.UserClient) *UserService {
	return &UserService{userClient: c}
}

// =========================
// CREATE USER
// =========================

func (s *UserService) CreateUser(ctx context.Context, req *userrpb.CreateUserRequest) (*userrpb.UserResponse, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("phone is required")
	}

	if len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	if req.Age <= 0 || req.Age > 120 {
		return nil, errors.New("invalid age")
	}

	if strings.TrimSpace(req.Address) == "" {
		return nil, errors.New("address is required")
	}

	return s.userClient.CreateUser(req)
}

// =========================
// GET BY ID
// =========================

func (s *UserService) GetUserByID(ctx context.Context, req *userrpb.GetUserByIDRequest) (*userrpb.UserResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.userClient.GetUserByID(req)
}

// =========================
// GET BY PHONE
// =========================

func (s *UserService) GetUserByPhone(ctx context.Context, req *userrpb.GetUserByPhoneRequest) (*userrpb.UserResponse, error) {

	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("phone is required")
	}

	if len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	return s.userClient.GetUserByPhone(req)
}

// =========================
// GET ALL USERS
// =========================

func (s *UserService) GetAllUsers(ctx context.Context, req *userrpb.GetAllUsersRequest) (*userrpb.GetAllUsersResponse, error) {

	return s.userClient.GetAllUsers(req)
}

// =========================
// UPDATE USER
// =========================

func (s *UserService) UpdateUser(ctx context.Context, req *userrpb.UpdateUserRequest) (*userrpb.UserResponse, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if req.Phone != "" && len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	if req.Age != 0 && (req.Age < 0 || req.Age > 120) {
		return nil, errors.New("invalid age")
	}

	return s.userClient.UpdateUser(req)
}

// =========================
// DELETE USER
// =========================

func (s *UserService) DeleteUser(ctx context.Context, req *userrpb.DeleteUserRequest) (*userrpb.Empty, error) {

	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	return s.userClient.DeleteUser(req)
}