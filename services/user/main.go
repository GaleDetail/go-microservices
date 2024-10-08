package main

import (
	"context"
	"github.com/GaleDetail/go-microservices/api/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	user, err := FetchUserData(req.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{Id: user.ID, Name: user.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	log.Println("User service is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
