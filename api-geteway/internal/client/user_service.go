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
	conn *grpc.ClientConn
	svc  userrpb.UserServiceClient
}

// =========================
// CONFIG
// =========================

const (
	timeoutUSER = 5 * time.Second
)

// =========================
// INIT (FAIL FAST)
// =========================

func NewUserClient(addr string) (*UserClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), tm)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // real connect wait
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ User service connected:", addr)

	return &UserClient{
		conn: conn,
		svc:  userrpb.NewUserServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *UserClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *UserClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// =========================
// METHODS
// =========================

func (c *UserClient) CreateUser(req *userrpb.CreateUserRequest) (*userrpb.UserResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.CreateUser(ctx, req)
}

func (c *UserClient) GetUserByID(req *userrpb.GetUserByIDRequest) (*userrpb.UserResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetUserByID(ctx, req)
}

func (c *UserClient) GetUserByPhone(req *userrpb.GetUserByPhoneRequest) (*userrpb.UserResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetUserByPhone(ctx, req)
}

func (c *UserClient) GetAllUsers(req *userrpb.GetAllUsersRequest) (*userrpb.GetAllUsersResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.GetAllUsers(ctx, req)
}

func (c *UserClient) UpdateUser(req *userrpb.UpdateUserRequest) (*userrpb.UserResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.UpdateUser(ctx, req)
}

func (c *UserClient) DeleteUser(req *userrpb.DeleteUserRequest) (*userrpb.Empty, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.svc.DeleteUser(ctx, req)
}