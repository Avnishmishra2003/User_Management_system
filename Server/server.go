package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "user_project/server/proto" // Replace with your generated package

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

// Create a new user
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := s.db.Exec(query, req.GetName(), req.GetEmail())
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Message: "User created successfully"}, nil
}

// Retrieve a user by ID
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := s.db.QueryRow(query, req.GetId())

	var id int32
	var name, email string
	err := row.Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

// Update a user
func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := s.db.Exec(query, req.GetName(), req.GetEmail(), req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Message: "User updated successfully"}, nil
}

// Delete a user
func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := s.db.Exec(query, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Message: "User deleted successfully"}, nil
}

func main() {
	// Connect to MySQL
	dsn := "username:password@tcp(localhost:3306)/grpc_project"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{db: db})

	fmt.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
