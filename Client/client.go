package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "user_project/server/proto" // Replace with your generated package

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Create a new user
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.CreateUser(ctx, &pb.CreateUserRequest{Name: "Avnish", Email: "Avm@example.com"})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	fmt.Println("User created successfully")
}
