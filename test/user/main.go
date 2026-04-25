package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/khbdev/what-food-proto/proto/userr"
)

var client userpb.UserServiceClient

func connect() {
	conn, err := grpc.Dial(
		"localhost:50050",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	client = userpb.NewUserServiceClient(conn)
}

// ================= CREATE =================
func createUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:    "Test",
		Phone:   "+998901234568",
		Age:     20,
		Address: "Tashkent",
		Email:   "test@gmail.com",
		Image:   "img.png",
	})
	if err != nil {
		log.Println("create error:", err)
		return
	}

	fmt.Println("CREATE:", res)
}

// ================= GET BY ID =================
func getByID() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GetUserByID(ctx, &userpb.GetUserByIDRequest{
		Id: 2,
	})
	if err != nil {
		log.Println("getByID error:", err)
		return
	}

	fmt.Println("GET BY ID:", res)
}

// ================= GET BY PHONE =================
func getByPhone() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GetUserByPhone(ctx, &userpb.GetUserByPhoneRequest{
		Phone: "+998911003630",
	})
	if err != nil {
		log.Println("getByPhone error:", err)
		return
	}

	fmt.Println("GET BY PHONE:", res)
}
	
// ================= GET ALL =================
func getAll() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
	if err != nil {
		log.Println("getAll error:", err)
		return
	}

	fmt.Println("GET ALL:", res)
}

// ================= UPDATE =================
func updateUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.UpdateUser(ctx, &userpb.UpdateUserRequest{
		Id:      1,
		Name:    "Updated",
		Phone:   "+998991112233",
		Age:     21,
		Address: "Tashkent",
		Email:   "updated@gmail.com",
		Image:   "new.png",
	})
	if err != nil {
		log.Println("update error:", err)
		return
	}

	fmt.Println("UPDATE:", res)
}

// ================= DELETE =================
func deleteUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.DeleteUser(ctx, &userpb.DeleteUserRequest{
		Id: 1,
	})
	if err != nil {
		log.Println("delete error:", err)
		return
	}

	fmt.Println("DELETE:", res)
}

// ================= MAIN =================
func main() {
	connect()

	// createUser()
	// getByID()
	// getByPhone()
	// getAll()
	// updateUser()
	// deleteUser()
}