package main

import (
	"context"
	"fmt"
	"log"
	"time"

	authpb "github.com/khbdev/what-food-proto/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connect() authpb.AuthServiceClient {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	return authpb.NewAuthServiceClient(conn)
}

func register(c authpb.AuthServiceClient, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	idStr := fmt.Sprintf("%d", id)

	phone := "+99845" + idStr + "434311"

	res, err := c.Register(ctx, &authpb.RegisterRequest{
		FullName: "Azizbek",
		Phone:    phone,
		Age:      21,
		Address:  "Tashkent",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("REGISTER:", res.Message)
}

func login(c authpb.AuthServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Login(ctx, &authpb.LoginRequest{
		Phone: "+99845455351",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("LOGIN:", res.Message)
}

func verify(c authpb.AuthServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.VerifyOTP(ctx, &authpb.VerifyRequest{
		Otp: 775855,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ACCESS:", res.AccessToken)
	log.Println("REFRESH:", res.RefreshToken)
}

func main() {
	client := connect()

	// 👉 o'zing tanlaysan
for i := 1; i < 50; i++ {
		register(client, i)
}
	// login(client)
	// verify(client)
}