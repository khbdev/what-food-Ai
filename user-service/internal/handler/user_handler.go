package handler

import (
	"context"

	userpb "github.com/khbdev/what-food-proto/proto/user;userpb"

	"user-service/internal/domain"
	"user-service/internal/models"
)

type UserHandler struct {
	userUC domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUC: uc,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.UserResponse, error) {
	err := h.userUC.Create(ctx, &models.CreateUserRequest{
		Name:  req.Name,
		Phone: req.Phone,
		Age:   int(req.Age),
		Email: req.Email,
		Image: req.Image,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.UserResponse, error) {
	user, err := h.userUC.GetByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		User: &userpb.User{
			Id:      uint64(user.ID),
			Name:    user.Name,
			Phone:   user.Phone,
			Age:     int32(user.Age),
			Address: user.Address,
			Email:   user.Email,
			Image:   user.Image,
		},
	}, nil
}

func (h *UserHandler) GetUserByPhone(ctx context.Context, req *userpb.GetUserByPhoneRequest) (*userpb.UserResponse, error) {
	user, err := h.userUC.GetByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		User: &userpb.User{
			Id:      uint64(user.ID),
			Name:    user.Name,
			Phone:   user.Phone,
			Age:     int32(user.Age),
			Address: user.Address,
			Email:   user.Email,
			Image:   user.Image,
		},
	}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	res, err := h.userUC.GetAll(ctx, &models.GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]*userpb.User, 0, len(res.Users))
	for _, u := range res.Users {
		users = append(users, &userpb.User{
			Id:      uint64(u.ID),
			Name:    u.Name,
			Phone:   u.Phone,
			Age:     int32(u.Age),
			Address: u.Address,
			Email:   u.Email,
			Image:   u.Image,
		})
	}

	return &userpb.GetAllUsersResponse{
		Users:  users,
		Total:  int32(res.Total),
		Limit:  int32(res.Limit),
		Offset: int32(res.Offset),
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UserResponse, error) {
	err := h.userUC.Update(ctx, uint(req.Id), &models.UpdateUserRequest{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.Empty, error) {
	err := h.userUC.Delete(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.Empty{}, nil
}