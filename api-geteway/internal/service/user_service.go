package service

import (
	"context"
	"errors"
	"strings"

	"api-geteway/internal/client"
	userrpb "github.com/khbdev/what-food-proto/proto/userr"
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

	if strings.TrimSpace(req.FullName) == "" {
		return nil, errors.New("full_name is required")
	}

	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("phone is required")
	}

	if len(req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	return s.userClient.CreateUser(req)
}

// =========================
// GET USER BY ID
// =========================

func (s *UserService) GetUserByID(ctx context.Context, req *userrpb.GetUserByIDRequest) (*userrpb.UserResponse, error) {

	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	return s.userClient.GetUserByID(req)
}

// =========================
// GET USER BY PHONE
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

	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	if req.FullName != nil && strings.TrimSpace(*req.FullName) == "" {
		return nil, errors.New("full_name cannot be empty")
	}

	if req.Phone != nil && len(*req.Phone) < 9 {
		return nil, errors.New("phone is invalid")
	}

	return s.userClient.UpdateUser(req)
}

// =========================
// DELETE USER
// =========================

func (s *UserService) DeleteUser(ctx context.Context, req *userrpb.DeleteUserRequest) (*userrpb.Empty, error) {

	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	return s.userClient.DeleteUser(req)
}