package main

import (
	"context"
	"log"
	"time"

	authpb "github.com/khbdev/what-food-proto/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// connect to auth service
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("connection error:", err)
	}
	defer conn.Close()

	client := authpb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// =========================
	// REGISTER TEST
	// =========================
	r1, err := client.Register(ctx, &authpb.RegisterRequest{
		FullName: "Azizbek",
		Phone:    "+998901112233",
		Age:      21,
		Address:  "Tashkent",
	})
	if err != nil {
		log.Fatal("register error:", err)
	}

	log.Println("REGISTER:", r1.Message)

	// =========================
	// LOGIN TEST
	// =========================
	r2, err := client.Login(ctx, &authpb.LoginRequest{
		Phone: "+998901112233",
	})
	if err != nil {
		log.Fatal("login error:", err)
	}

	log.Println("LOGIN:", r2.Message)

	// =========================
	// VERIFY TEST (manual OTP)
	// =========================
	r3, err := client.VerifyOTP(ctx, &authpb.VerifyRequest{
		Otp: 123456, // Redisdan chiqqan OTP ni qo'yasan
	})
	if err != nil {
		log.Fatal("verify error:", err)
	}

	log.Println("ACCESS TOKEN:", r3.AccessToken)
	log.Println("REFRESH TOKEN:", r3.RefreshToken)
}