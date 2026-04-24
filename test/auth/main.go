package main

import (
	"context"
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

func register(c authpb.AuthServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Register(ctx, &authpb.RegisterRequest{
		FullName: "Azizbek",
		Phone:    "+99845455351",
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
		Phone: "+998901112233",
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
		Otp: 688779,
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
	register(client)
	// login(client)
	// verify(client)
}