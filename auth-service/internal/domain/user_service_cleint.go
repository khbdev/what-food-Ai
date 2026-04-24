package domain


import userrpb "github.com/khbdev/what-food-proto/proto/userrpb"

type UserService interface {
	CreateUser(req *userrpb.CreateUserRequest) (*userrpb.UserResponse, error)
	GetUserByPhone(req *userrpb.GetUserByPhoneRequest) (*userrpb.UserResponse, error)
}