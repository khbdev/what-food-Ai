package handler

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/models"

	userrpb "github.com/khbdev/what-food-proto/proto/userr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userrpb.UnimplementedUserServiceServer
	usecase domain.UserUsecase
}

func NewUserHandler(usecase domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userrpb.CreateUserRequest) (*userrpb.UserResponse, error) {
	err := h.usecase.Create(ctx, &models.CreateUserRequest{
		Name:    req.Name,
		Phone:   req.Phone,
		A:     int(req.Age),
		Address: req.Address,
		Email:   req.Email,
		Image:   req.Image,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userrpb.UserResponse{}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *userrpb.GetUserByIDRequest) (*userrpb.UserResponse, error) {
	user, err := h.usecase.GetByID(ctx, uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &userrpb.UserResponse{User: toProto(user)}, nil
}

func (h *UserHandler) GetUserByPhone(ctx context.Context, req *userrpb.GetUserByPhoneRequest) (*userrpb.UserResponse, error) {
	user, err := h.usecase.GetByPhone(ctx, req.Phone)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &userrpb.UserResponse{User: toProto(user)}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *userrpb.GetAllUsersRequest) (*userrpb.GetAllUsersResponse, error) {
	result, err := h.usecase.GetAll(ctx, &models.GetAllUsersRequest{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUsers := make([]*userrpb.User, 0, len(result.Users))
	for i := range result.Users {
		pbUsers = append(pbUsers, toProto(&result.Users[i]))
	}

	return &userrpb.GetAllUsersResponse{
		Users:  pbUsers,
		Total:  int32(result.Total),
		Limit:  int32(result.Limit),
		Offset: int32(result.Offset),
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *userrpb.UpdateUserRequest) (*userrpb.UserResponse, error) {
	err := h.usecase.Update(ctx, uint(req.Id), &models.UpdateUserRequest{
		Name:    req.Name,
		Phone:   req.Phone,
		Age:     int(req.Age),
		Address: req.Address,
		Email:   req.Email,
		Image:   req.Image,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userrpb.UserResponse{}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *userrpb.DeleteUserRequest) (*userrpb.Empty, error) {
	if err := h.usecase.Delete(ctx, uint(req.Id)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userrpb.Empty{}, nil
}

// toProto — models.User -> userrpb.User
func toProto(u *models.User) *userrpb.User {
	return &userrpb.User{
		Id:        uint64(u.ID),
		Name:      u.Name,
		Phone:     u.Phone,
		Age:       int32(u.Age),
		Address:   u.Address,
		Email:     u.Email,
		Image:     u.Image,
		Role:      userrpb.Role(u.Role),
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}