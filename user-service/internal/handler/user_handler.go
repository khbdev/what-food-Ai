package handler

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/models"

	userrpb package jwt

import (
	"auth-service/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateRefreshToken(m models.TokenModel) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")


	expStr := os.Getenv("JWT_REFRESH_EXP_DAYS")
	if expStr == "" {
		expStr = "30"
	}

	expDays, err := time.ParseDuration(expStr + "24h")
	if err != nil {
		expDays = 30 * 24 * time.Hour
	}

	claims := RefreshClaims{
		UserID: m.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDays)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
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
	err := h.usecase.Create(ctx, &models.User{
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
	result, err := h.usecase.GetAll(ctx, &models.GetAllUsersRequest{})
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
	err := h.usecase.Update(ctx, uint(req.Id), &models.User{
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

func mapRole(r models.Role) userrpb.Role {
	switch r {
	case models.RoleUser:
		return userrpb.Role_ROLE_USER
	case models.RoleAdmin:
		return userrpb.Role_ROLE_ADMIN
	default:
		return userrpb.Role_ROLE_UNSPECIFIED
	}
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
		Role:      mapRole(u.Role),
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}