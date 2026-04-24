package client

import (
	"context"
	"log"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userrpb "github.com/khbdev/what-food-proto/proto/userr"
)

type UserClient struct {
	conn   *grpc.ClientConn
	client userrpb.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	log.Fatal("user service")
	return &UserClient{
		conn:   conn,
		client: userrpb.NewUserServiceClient(conn),
	}, nil
}

// CreateUser
func (c *UserClient) CreateUser(req *userrpb.CreateUserRequest) (*userrpb.UserResponse, error) {
	var res *userrpb.UserResponse
	var err error

	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		res, err = c.client.CreateUser(ctx, req)
		cancel()

		if err == nil {
			return res, nil
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil, err
}

// GetUserByPhone
func (c *UserClient) GetUserByPhone(req *userrpb.GetUserByPhoneRequest) (*userrpb.UserResponse, error) {
	var res *userrpb.UserResponse
	var err error

	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		res, err = c.client.GetUserByPhone(ctx, req)
		cancel()

		if err == nil {
			return res, nil
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil, err
}

func (c *UserClient) Close() error {
	return c.conn.Close()
}